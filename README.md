## Golang Domain Driven Design

Requires: Postgres (edit `RepositoryOpts.Addr` struct in `cmd/main.go` and `repo/pg_test.go`)
Build: `cd cmd && go build && ./cmd`
Test: `go vet ./... && go test ./...`
Address: `http://localhost:8080/`

## API

### `POST /signup`
**Request**:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email": "jackson@juandefu.ca","password": "pass","firstName": "Jackson","lastName": "Sabey"}' \
  http://localhost:8080/signup
```
**Body**:
```json
{
  "email": "jackson@juandefu.ca",
  "password": "pass",
  "firstName": "Jackson",
  "lastName": "Sabey"
}
```

**Response**:
```json
{
  "token": "jwt-token" 
}
```

### `POST /login`
**Request**:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"email": "jackson@juandefu.ca","password": "pass"}' \
  http://localhost:8080/login
```
**Body**:
```json
{
  "email": "jackson@juandefu.ca",
  "password": "pass"
}
```

**Response**:
```json
{
  "token": "jwt-token"
}
```

### `GET /users`
**Request**:
```
curl --header "X-Authentication-Token: jwt-token" \
  http://localhost:8080/users
```
**Response**:
```json
{
  "users": [
    {
      "email": "jackson@juandefu.ca",
      "firstName": "Jackson",
      "lastName": "Sabey"
    }
  ]
}
```

### `PUT /users`
**Request**:
```
curl --header "X-Authentication-Token: jwt-token" --header "Content-Type: application/json" \
  --request PUT \
  --data '{"firstName": "JACKSON","lastName": "SABEY"}' \
  http://localhost:8080/users
```
```json
{
  "firstName": "JACKSON",
  "lastName": "SABEY"
}
```

**Response**:
`none`
