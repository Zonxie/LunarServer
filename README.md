# RESTful server

This service provides the outfacing REST api that users contact to make changes to their account.
Runs on port: 5000

### First time setup

```
go get ./...
```

### Running the application

```
go run .\server\cmd\server.go
```

### Secret Token needed to contact this service
Authorization: Bearer SuperSecretAuth


### Running Tests

Install REST Client extension and test through the test.http file


### TODO list / What could be different

[] Move each service into docker containers
[] Change communication between server and engine too messaging