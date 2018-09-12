
## Go Rest API application with clean architecture


This repo contains code for basic Rest Application with CRUD operations. PostgresQL database is used to implement the application.

### Following features are implemented:

* JWT Authentication
* Mux Rest
* [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
* Postgres Connectivity
* Externalized Configuration (Default properties provided through configuration.properties can be overwritten through environment variables)
* Formatted Rest Error Response (Status Code and Messages)
* Unit Testing

##### To-DO:

* Complete pending CRUD operations
* Implement Flags for Configuration


##### Database setup

create database contactsdb;

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(200) NOT NULL,
    lastname VARCHAR(200) NOT NULL,
    address VARCHAR(200)
);


CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    displayname VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL
);

##### Testing the application

* Navigate to the folder with code and start the application by running following command
	go run main.go
* create user by posting the below data to http://localhost:8000/create-user
	{"Username" : "admin", "Displayname" : "admin", "Password" : "admin@123"}
* Login to the application by making a post call with below data to http://localhost:8000/login
	{"Username" : "admin", "Displayname" : "admin", "Password" : "admin@123"}
* From the response capture the Token
* To create contacts set the Token obtained in previous step as "Bearer Token" in Authorization
* Make post call http://localhost:8000/create-contact with following data
	{"firstname":"first","lastname":"contact"}

