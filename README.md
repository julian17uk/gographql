# README #

Graphql is a GraphQL server written in go (golang). 


### What is this repository for? ###

* This project has been written by Julian Karnik at ECS Digital. When the graphql server is running it allows users to access and mutate the data using any application designed to interact with a graphql server. For example the application GraphiQL can be used to test the project.
* Version 1.0

### How do I get set up? ###

* First, please ensure that golang is available. See https://golang.org/doc/install
* Configuration: none required
* Dependencies
* cmd folder hold the main.go file
* internal folder holds the internal functions
* How to run tests: user$ go test -v (from within the test folder)
* The data is stored in the data folder

### How to run graphql server in the terminal? ###

* Run, from cmd folder user$ go run main.go
* Using GraphiQL (or others) make a call to the server, for example. To install GraphiQL see 


### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* For further information please contact Julian Karnik at julian.karnik@ecs-digital.co.uk
