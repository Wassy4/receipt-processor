# Instructions

1. Run `make build` to build the Docker image.
2. Run `make run` to run the Docker image as a container.
    - API calls can be issued through `http://localhost:8080/`
3. Run `make stop` to stop the container.
4. Run `make delete` to delete the container.

# Usage example

## `POST /receipts/process`

### Request
```
curl -i -X POST -H "Content-Type: application/json" http://localhost:8080/receipts/process --data '{request body}'
```

### Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 02 Dec 2024 01:09:50 GMT
Content-Length: 51

{
  "id": "9b0030c3-1353-4b02-9e3c-644e"
}
```

## `GET /receipts/{id}/points`

### Request
```
curl http://localhost:8080/receipts/9b0030c3-1353-4b02-9e3c-644e/points
```

### Response
```
{
  "points": 28
}
```