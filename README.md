# README #

This project is a GraphQL server written in go (Golang). GraphQL is a query language for APIs, this represents a significant improvement over REST APIs. Please refer to https://graphql.org/ for further details.


### What is this repository for? ###

* This project has been written by Julian Karnik at ECS Digital. When the GraphQL server is running it allows users to access and mutate the data using any application designed to interact with a GraphQL API. For example the application GraphiQL can be used to test the API.
* Version 1.0

### How do I get set up? ###

* To set up first install golang on your device. See https://golang.org/doc/install
* To install GraphiQL see https://www.electronjs.org/apps/graphiql (not required)
* Configuration: none required
* Dependencies: This project uses the graphql go library. See https://github.com/graphql-go/graphql (to install run go get github.com/graphql-go/graphql) 
* cmd folder hold the main.go file
* internal folder holds the internal functions
* How to run tests: user$ go test -v (from within the test folder)
* The data is stored in the data folder

### How to run graphql server in the terminal? ###

* To launch the API, from cmd folder run user$ go run main.go
* To test the API, run GraphiQL (or others) to make a call to the API.
* When the service is running the API Endpoint is http:xxxxxxxx
* Example GraphQL API calls are available in the examples folder.
* There are two basic types of API calls get (same as GET in REST) and mutation (this covers POST, PUT & DELETE in REST)

### Warning ###

* When this service is running, it will accept http connections on port 8080 from anywhere on the internet
* This service allows any other computer to get & mutate files stored on this computer, therefore please use cautiously


### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* For comments contact julian.karnik@ecs-digital.co.uk
