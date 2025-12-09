# frogsmash
Compare frogs and other things

# Docker 
Services can be run from `docker-compose.yml`. This can be done through the VS Code UI or by running `docker compose up -d --build`

pgAdmin is included in the compose to allow direct database querying. To connect to the db, run the services, go to `http://localhost:5050`, input the email and password from `.env` to log in. Then, right-click servers, register, server... and add the connection details. Host name will be `db` (that's what the db service is called in `docker-compose.yml`). Port is the standard 5432 (also in `docker-compose`). Then, maintenance database is whatever the database is called in `.env`, and username and password are what they are in the connection string in `.env`.

There are two web services, a prod one that uses npm build and nginx, and a dev one that simply runs `npm run dev`. You'll generally use the dev one. 

# Migration scripts: 

`migrate create -ext sql -dir db/migrations <migration_name>` creates empty migration scripts. Migration automatically gets run at API startup. This might be something we change in the future. 

# Swagger: 
Run `swag init -g .\cmd\server\main.go` to update swagger docs. Make sure to annotate endpoints and dtos so they show up correctly. Running this may cause some invalid delims to appear in the docs file, just delete those. I don't know what to do about that yet. 

# Fly.io:
Deploy an instance like thus from the api folder: `fly deploy --config fly.server.toml`