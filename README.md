# test

Simple CR(U)D Go microservice used for testing read-only API mode

## Usage

- `make up` - run service in Docker container

- `make down` - stop service

## API

### Models

- Item
```json
{
    "id": 1, // only in resp body
    "title": "Walk the dog",
    "description": "Take the dog for a 30-minute walk",
    "done": false
}
```

- User
```json
{
    "id": 1, // only in resp body
    "name": "ed",
    "job": "backend dev"
}
```

- API Status
```json
{
    "table_name": "users",
    "read_only": false
}
```

### Routes

- [GET]     `/api/items`
- [POST]    `/api/items` (Item model as req body)
- [GET]     `/api/items/{id}`
- [DELETE]  `/api/items/{id}`

- [GET]     `/api/users`
- [POST]    `/api/users` (User model as req body)
- [GET]     `/api/users/{id}`
- [DELETE]  `/api/users/{id}`

- [GET]     `/api/status`
- [POST]    `/api/status/{read_only}` (true/false as query param)
