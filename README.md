# frogsmash
Compare frogs and other things

# Docker 
Services can be run from `docker-compose.yml`. This can be done through the VS Code UI or by running `docker compose up -d --build`

pgAdmin is included in the compose to allow direct database querying. To connect to the db, run the services, go to `http://localhost:5050`, input the email and password from `.env` to log in. Then, right-click servers, register, server... and add the connection details. Host name will be `db` (that's what the db service is called in `docker-compose.yml`). Port is the standard 5432 (also in `docker-compose`). Then, maintenance database is whatever the database is called in `.env`, and username and password are what they are in the connection string in `.env`.
