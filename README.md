# GoLang microservices system using gRPC and MongoDB

## Basic Information

This is a POC dummy login system using goLang, gRPC, MongoDB, Docker in a microservices'
architecture.

In MongoDB we store some users. The client does a /login call to the proxy service, 
which contacts with the db-handler service with gRPC to verify and get
the user's info. Then the proxy creates a jwt token with the user's info and returns 
the token to the client. Client then can call the /info endpoint to see it's logged 
in user's info.

Note: The MongoDB `users` collection is filled with a single User with credentials 
`username:password` in the beginning of the App with a go routine

## Endpoints

Endpoints:
- `/login` requires username and password and returns a JWT token containing some random user info stored in the DB, if the info is valid 
- `/info` uses the jwt token's claims to show the user the info stored in db

## Run

1. Build by running `docker-compose up`
1. Make `POST` call in `localhost:8000/login` with `{"username":"username","password":"password"}` JSON body
1. Make `GET` call `localhost:8000/info`