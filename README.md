# Companies REST API Project

This is a Companies GoLang REST API project that manages company data. It provides endpoints to retrieve, create, update, and delete company information.

## Prerequisites

- GoLang 1.21.1
- Docker (for running the PostgreSQL database)

## Setting Up the Database

1. Navigate to the `/db` directory:

   ```bash
   cd /db
   ```

2. Run the following command to start the PostgreSQL database in a Docker container with initial data:

```bash 
docker-compose up -d
```
This will create a PostgreSQL database with the specified schema and initial data.

Note: Ensure Docker is installed and running on your system before executing this command.

## Running the Application
Make sure you're in the root directory of the project.

Run the following command to start the application and expose it on port 8080:

```bash
docker-compose up -d
```
This will build and run the GoLang application within a Docker container.

## Configuration
Configuration for the application is set up in the configs/config.yaml file. You can modify the following database details in this file:

- host: The hostname or IP address of the PostgreSQL database server.
- port: The port on which the PostgreSQL database server is running.
- user: The PostgreSQL username.
- password: The PostgreSQL password.
- dbname: The name of the PostgreSQL database.

## Endpoints
The API provides the following endpoints:

- GET /company/:id: Retrieves a company by its ID.
- POST /company: Creates a new company.
- PATCH /company: Updates an existing company.
- DELETE /company/:id: Deletes a company by its ID.

### Example Usage
To retrieve a company by ID:

```bash
curl http://localhost:8080/company/{id}
```

To create a new company (provide JSON data in the request body):

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Company D",
    "description": "Description for Company D",
    "amountOfEmployees": 10,
    "registered": true,
    "companyType": "Sole Proprietorship"
}' http://localhost:8080/company
```

To update an existing company (provide JSON data in the request body):

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{
    "name": "Company D",
    "description": "Description for Company D",
    "amountOfEmployees": 10,
    "registered": true,
    "companyType": "Sole Proprietorship"
}' http://localhost:8080/company
```
To delete a company by ID:

```bash
curl -X DELETE http://localhost:8080/company/{id}
```

Please replace {id} with the actual company ID when making requests.

## Requirements

Technical requirements
Build a microservice to handle companies. It should provide the following operations:
- Create
- Patch
- Delete
- Get (one)

Each company is defined by the following attributes:
- ID (uuid) required
- Name (15 characters) required - unique
- Description (3000 characters) optional
- Amount of Employees (int) required
- Registered (boolean) required
- Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) required

Only authenticated users should have access to create, update and delete companies.

Expectations:
As a deliverable, we expect a GitHub repository (or any other git based repo) with the source
code. We would like the solution to contain clear instructions to set up and execute the project.
We expect the solution to be production ready.

Will be considered a plus:
- On each mutating operation, an event should be produced.
- Dockerize the application to be ready for building the production docker image
- Use docker for setting up the external services such as the database
- REST is suggested, but GRPC is also an option
- JWT for authentication
- Kafka for events
- DB is up to you
- Integration tests are highly appreciated
- Linter
- Configuration file