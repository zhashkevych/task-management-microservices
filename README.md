# Microservices Structure Example (Task Management Software)
## Stack
- Go
- Postgres
- <a href="https://www.krakend.io/">KrakenD</a> as API Gateway
- <a href="https://github.com/lpereira/lwan"> Lwan</a> (as file server)
- Docker & Docker-Compose

## Requirements
- Installed Docker
- Installed `golang-migrate` 

## CLI Tools
- `make help` to see all available instructions
- `make run` to build & run project
- `make migrate-users` & `make migrate-tasks` to apply database migrations

## TODO
- Inter Process Communication (gRPC / Message Queue)
- Logging & Tracing
