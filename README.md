# E-commerce API

## Overview

This project is a simple e-commerce RESTful API, designed as part of an assessment. 
The API handles basic CRUD operations for products and orders, and includes user management and authentication. 
The project is intentionally simplified and **not production-ready**, 
as certain security best practices have been intentionally omitted to focus on core functionality 
as specified by the assessment instructions.

## Features

- **User Management**: Register new users, login to receive JSON Web Tokens (JWT), and authenticate sessions.
- **Product Management**: Admin-only access to create, read, update, and delete products.
- **Order Management**: Place and manage orders, with the ability to cancel pending orders and update order status (admin privilege).

## Technical Stack

- **Language**: Go (Golang)
- **Framework**: chi for routing
- **Authentication**: JWT for session management
- **Database**: PostgreSQL
- **GORM**: Since the assessment instruction didn't specify a data access preference (SQL queries / ORM ),
            an ORM (GORM) was used because it is more ideal for the project purpose
- **Documentation**: Swagger for endpoint documentation

## Endpoints
Find details of the implemented endpoint in [documentation](docs/swagger.yaml)

## Installation and Deploying

1. **Clone the repository**
   ```bash
   git clone https://github.com/ercross/instashop.git
   cd instashop
   ```
2. Find the configuration files (i.e., Dockerfile, compose.yml, and Makefile) 
   provided to automatically deploy the containerized app locally through docker compose.
3. Run `make deploy`
4. Query the app at `localhost:15001` 

## Project Limitations

This project is intended for demonstration purposes and is not production-ready. **Security limitations** include:

- Simplified endpoint protections to focus on assessment requirements.
- No advanced security features like rate limiting, request throttling, or complex role-based access controls.
- Basic JWT token handling without refresh tokens or secure storage recommendations.

These limitations are intentional, as the focus of this assessment is on demonstrating basic functionality, not on production-grade security.

## Future Improvements

For production readiness, the following improvements are recommended:

- **Enhanced Security**: Add rate limiting, CORS, HTTP/1.1 security headers, stricter input validation, and secure JWT handling.
- **Role-based Access Control**: Implement a more robust access control system to handle different user roles more flexibly.
- **Logging**: Implement logging to improve error-tracing, debugging and troubleshooting
- **Performance**: Implement pagination to increase page load time
- **Integrated Tests**: implement necessary handler tests to improve stability

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.