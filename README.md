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

## Build & Run Project
- Run `make run` in terminal window
- **If you are running project locally for the first time,** open new window and run `make migrate-users && make migrate-tasks`

## TODO
- Inter Process Communication (gRPC / Message Queue)
- Logging & Tracing

## Endpoints
### POST /user/sign-up

Create New User

**Example Input**:
```json
{
  "first_name": "Albert",
  "last_name": "Einstein",
  "username": "generalrelativity",
  "password": "E=mc2baby"
}
```
**Example Response**:
```json
{
  "id": 1
}
```

### GET /user/token

Generate JWT Token

**Example Input**:
```json
{
  "username": "generalrelativity",
  "password": "E=mc2baby"
}
```
**Example Response**:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsImtpZCI6IjEifQ.eyJhdWQiOiJodHRwOi8vZ2F0ZXdheTo4MDgwIiwiZXhwIjoxNjAzMzAzMzUyLCJpc3MiOiJodHRwOi8vdXNlcnMtc2VydmljZTo4MDAwIiwidXNlcl9pZCI6MX0.nL3ZtfqxBsgMLFcPrX16MekQWrduWE3dAUGFhm1bZzI"
}
```

### GET /user/profile

Return User Info

**HTTP Headers**:

```
Authorization: Bearer <access_token>
```

**Example Response**:
```json
{
  "id": 1,
  "first_name": "Albert",
  "last_name": "Einstein",
  "username": "generalrelativity",
  "password": "E=mc2baby"
}
```

### POST /tasks

Create New Task

**HTTP Headers**:

```
Authorization: Bearer <access_token>
```

**Example Input**:
```json
{
  "title": "Task 1"
}
```
**Example Response**:
```json
{
  "id": 1
}
```

### GET /tasks

Get All Tasks

**HTTP Headers**:

```
Authorization: Bearer <access_token>
```

**Example Response**:
```json
{
  "tasks": [
    {
      "id": 1,
      "title": "Task 1",
      "created_at": "2020-10-21T06:09:26.654986Z",
      "user_id": 1
    }
  ]
}
```

### GET /tasks/1

Get Task By Id

**HTTP Headers**:

```
Authorization: Bearer <access_token>
```

**Example Response**:
```json
{
  "id": 1,
  "title": "Task 1",
  "created_at": "2020-10-21T06:09:26.654986Z",
  "user_id": 1
}
```