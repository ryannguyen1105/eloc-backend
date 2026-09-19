# Dien May Loc - Microservices E-Commerce Backend

A microservices-based backend system for a consumer electronics and home appliances e-commerce retail store, focusing on scalable system design and modular architecture.

---

## System Architecture

This project adopts the Database-per-Service pattern to maintain tight domain boundaries and service independence across 3 isolated PostgreSQL database instances:

- Auth Service (auth): Manages user identity, authentication, and token generation.
- Product Service (product): Handles consumer electronics catalog, categories, and inventory.
- Order Service (order): Manages shopping carts, checkout lifecycles, and transaction processing.

---

## Tech Stack

- Languages: Golang, SQL
- Frameworks & Libraries: Gin Gonic, SQLC, Viper, Testify
- Databases: PostgreSQL (Neon DB)
- API & Security: RESTful APIs, gRPC, Protocol Buffers, JWT, PASETO, Bcrypt
- DevOps & Infrastructure: Docker, GitHub Actions, Render, SSL, Git, Postman

---

## Features Implemented

### 1. Microservices Architecture & Data Access
- Applied Database-per-Service pattern with 3 PostgreSQL instances (auth, product, order).
- Integrated SQLC for compile-time type-safe Go code generation, eliminating boilerplate and reflection overhead.
- Implemented ACID transactions for concurrent write consistency during critical operations.

### 2. REST APIs & Security
- Built RESTful APIs using Gin framework with standardized JSON responses and centralized error handling middleware.
- Secured endpoints using Bcrypt password hashing, JWT stateless authentication, and custom middleware for token validation.

### 3. Configuration & Containerization
- Centralized environment management with Viper following 12-Factor App principles.
- Optimized Docker image sizes using Multi-stage builds with Alpine Linux for lightweight deployment.

### 4. CI/CD & Cloud Deployment
- Automated linting and unit testing (testify/require) via GitHub Actions CI workflows.
- Deployed Go microservices to Render connected to Neon PostgreSQL over SSL.

---

## Upcoming Roadmap

- Inter-Service Communication & Caching: Implementing gRPC/Protobuf protocols for internal service calls and integrating Redis for high-performance caching.
- Data Layer & Access Control: Migrating database drivers to pgx for enhanced PostgreSQL performance and implementing RBAC (Role-Based Access Control).
- Security & Observability: Configuring CORS middleware for secure cross-origin requests and setting up API Gateway with centralized observability.