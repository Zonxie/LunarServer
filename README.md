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