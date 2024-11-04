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
- **Documentation**: Swagger for endpoint documentation

## Endpoints

The following endpoints are included in the API:

### User Management / Authentication
- `POST /register`: Register a new user with email and password
- `POST /login`: Login and receive a JWT for session management

### Product Management (Admin-only access)
- `POST /products`: Create a new product
- `GET /products`: Retrieve a list of products
- `GET /products/{id}`: Retrieve a specific product
- `PUT /products/{id}`: Update an existing product
- `DELETE /products/{id}`: Delete a product

### Order Management (Authenticated users)
- `POST /orders`: Place an order for one or more products
- `GET /orders`: List all orders for the authenticated user
- `GET /orders/{id}`: View details of a specific order
- `PUT /orders/{id}/cancel`: Cancel an order if it is still pending
- (Admin-only) `PUT /orders/{id}/status`: Update the status of an order

## Installation and Deploying

1. **Clone the repository**
   ```bash
   git clone https://github.com/ercross/instashop.git
   cd ecommerce-api
   ```

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
- **Improved Error Handling**: Include more specific error messages.
- **Logging**: Implement logging to improve error-tracing, debugging and troubleshooting

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.