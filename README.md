# FrogSmash
FrogSmash is an app where you can compare different things. On the `smash` page, you are given two random items and you have to decide between them which you prefer. The criteria for this decision remain ambiguous. Items are created by a dev and can be images they have found on the internet (with proper licensing), or an image that a user has uploaded. 

## Technical details 
Below is a summary of the tech involved in FrogSmash

### Golang backend 
The backend is written in Go and consists of a number of deployed services.

#### API
The api is built with Gin and deployed on fly.io in production. The entrypoint to the api is `cmd\server\main.go`. This calls functions to set up the config, configure the container (dependency injection), and set up the http server. 

The api endpoint definitions and everything related to receiving requests and sending responses live in the `delivery` directory. Here we set up the middleware pipeline (which includes a Redis based rate-limiter for all endpoints, and auth filters for certain endpoints to check whether users are logged in and verified).

#### Score updater
The score updater is a separately deployed service that runs with a configured frequency. It asynchronously checks the `events` table in the database, which acts as a queue of comparison events to be processed, and updates the scores of items in the `items` table. Currently, it's set up to only run one score updater at a time, otherwise the order of event processing would not be guaranteed, which would affect the scores items receive. A future improvement would be to allow multiple events to be processed at once if the items involved don't overlap (e.g. A vs B could be handled at the same time as C vs D because the scores of one don't affect the scores of the other, but A vs B could not be handled at the same time as B vs C).

Scores of items are calculated using the Elo formula with a configurable k factor. This is the same formula used in chess to calculate the ratings of players following matches. It essentially weights the change in score of a given item by the difference in the two items' scores in a match up. E.g. if an item with score 500 beat an item with score 1500, the former would gain a lot of points and the latter would lose a lot. If, however, an item with a score of 1000 beat another with 1000, the change in scores would be much smaller. 

#### Migrator 
This is a deployable service that runs migrations on the configured database.

#### Worker 
The worker is a background process that currently only handles sending out emails asynchronously. In the system we have both a message producer and a message consumer both tied to a Redis database. The worker handles the consumption side, listening to messages sent out by other services. This setup means that flows with side effects don't depend on the side effects. For example, registration has a side effect of sending out a verification email, but the registration endpoint doesn't care whether the email has in fact been sent out. Instead, the registration service queues a message that the worker picks up in its own time and dispatches to the verification service to handle. 

#### Backend architecture 
The backend is designed with a layered architecture in mind. The philosophy is fundamentally based on a separation of concerns. Each layer has its own responsibilities, and nothing should fail in a layer due to an external responsibility. 

##### Delivery layer 
As mentioned previously, the delivery directory acts as an architectural layer. It is the gateway to the world outside of the server. It ensures that anything coming in (requests) and going out (responses) is properly formatted. It doesn't handle any of the business logic, rather it hands that work over to the app layer. 

##### App layer 
The app layer handles business logic. It consists of services. Services orchestrate behaviour. For example, the auth service has a register function which calls a hasher to hash the password it receives, calls a user service to create a user, and calls a message service to queue a registration event. Crucially, the auth service doesn't care about how passwords are hashed, how users are created or any side effects of registration, it just makes sure that the steps are carried out and any errors are bubbled up. The presentation of errors is handled by the delivery layer. 

##### Data layer 
The data layer handles data persistence. These are the concrete details of where data is stored, for example users, events, images and items. The service layer depends on the data layer via repos and abstract clients. Most of the data is stored in a postgres database. Images are stored in a Cloudflare R2 bucket and messages are stored in Redis, but nothing in the service layer is aware of this. 

### React frontend
The frontend is a static SPA built with React and styled with Material UI. Just like the backend, it is split into layers. 

#### App 
The app directory handles global configuration, providers and setting up routes. It is also the entrypoint into the site. 

#### Shared 
This is where pages, components, and hooks that don't belong to a specific feature or are referenced by multiple features live. 

#### Features 
This is where all the self-contained features live. It is organised in the same way that the features in the backend are. The intention is that each feature should be individually deployable, such that one doesn't depend on another. 

#### Componentisation 
Components roughly follow the principles of atomic design. They are divided based on complexity. Pages are the most complex and atoms the least. 

## Dev reference
Below are some useful scripts and things to know from a dev perspective

### Docker 
Services can be run from `docker-compose.yml`. This can be done through the VS Code UI or by running `docker compose up -d --build`

pgAdmin is included in the compose to allow direct database querying. To connect to the db, run the services, go to `http://localhost:5050`, input the email and password from `.env` to log in. Then, right-click servers, register, server... and add the connection details. Host name will be `db` (that's what the db service is called in `docker-compose.yml`). Port is the standard 5432 (also in `docker-compose`). Then, maintenance database is whatever the database is called in `.env`, and username and password are what they are in the connection string in `.env`.

There are two web services, a prod one that uses npm build and nginx, and a dev one that simply runs `npm run dev`. You'll generally use the dev one. 

### Migration scripts: 

`migrate create -ext sql -dir db/migrations <migration_name>` creates empty migration scripts. Migration automatically gets run at API startup. This might be something we change in the future. 

### Swagger: 
Run `swag init -g .\cmd\server\main.go` to update swagger docs. Make sure to annotate endpoints and dtos so they show up correctly. Running this may cause some invalid delims to appear in the docs file, just delete those. I don't know what to do about that yet. 

### Fly.io:
Deploy an instance like thus from the api folder: `fly deploy --config fly.server.toml`