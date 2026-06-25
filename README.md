# Dien May Loc Backend

A microservices-based backend system for a consumer electronics and home appliances e-commerce retail store, focusing on scalable system design.

## System Architecture

The project utilizes a Database-per-Service pattern with 3 isolated PostgreSQL instances to ensure service independence:

* Auth Service (auth): Handles user identity management and access tokens.
* Product Service (product): Manages the consumer electronics catalog and inventory.
* Order Service (order): Processes shopping carts and checkout lifecycles.

## Tech Stack

* Languages: Golang, SQL
* Frameworks & Libraries: Gin Gonic, SQLC, Testify (Unit Testing)
* Databases: PostgreSQL
* API & Security: RESTful APIs, gRPC, Protocol Buffers, JWT, PASETO, Bcrypt
* DevOps & Infrastructure: Docker, Git, Postman, GitHub Actions

## Features Implemented

### 1. Database Architecture & Transaction Management
* Designed a Database-per-Service pattern with 3 isolated PostgreSQL instances (auth, product, order).
* Implemented ACID transactions in auth and order services to ensure data consistency during concurrent writes.

### 2. Data Access Layer
* Integrated SQLC to generate type-safe Go code from raw SQL, optimizing performance and eliminating boilerplate.

### 3. RESTful APIs
* Implemented REST APIs using Gin Gonic with standardized JSON responses and error handling.

### 4. DevOps & Testing
* Containerized databases via Docker.
* Built automated CI pipelines using GitHub Actions to run linters and testify unit tests on every Pull Request.

## Upcoming Roadmap

* Production Readiness: Completing the remaining e-commerce core services and business logic to prepare the entire application for production.
* Cloud Deployment: Designing containerized deployment workflows to migrate the microservices infrastructure onto a cloud platform.
