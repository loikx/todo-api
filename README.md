# TODO

## Simple json api app
* /api/todo/create
* /api/todo/update
* /api/todo/find

## Stack
* docker
* docker-compose
* golang
* postgresql

### [POST] Create request /api/todo/create
* request
```json
{
  "name": "todo name",
  "body": "todo body",
  "priority": "low"
}
```

* response 
```json
{
  "id": "uuid",
  "name": "todo name",
  "body": "todo body",
  "priority": "low",
  "deadline": null,
  "createdAt": "some time",
  "updatedAt": "somt time"
}
```

### [POST] Update request /api/todo/update
* request
```json
{
  "id": "uuid",
  "name": "new todo name",
  "body": "new todo body",
  "priority": "high"
}
```

### [POST] Find request /api/todo/find
* request
```json
{
  "id": "uuid"
}
```

* response
```json
{
  "id": "uuid",
  "name": "todo name",
  "body": "todo body",
  "priority": "low",
  "deadline": null,
  "createdAt": "some time",
  "updatedAt": "somt time"
}
```
