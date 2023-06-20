# Pizza API

## Overview

This API allows users to order pizzas from a menu, track their orders, and rate their experiences. Users need to register and log in to use the API. The API uses JSON Web Tokens (JWT) for authentication and authorization.

## Base URL

The base URL for the API is `https://pizza.example.com/api/v1`.

## Authentication

To authenticate with the API, users need to send a `POST` request to the `/auth/login` endpoint with their email and password in the request body. The API will return a JWT in the response body if the credentials are valid. The JWT should be included in the `Authorization` header of subsequent requests as a bearer token.

Example request:

```http
POST /auth/login HTTP/1.1
Host: pizza.example.com
Content-Type: application/json

{
  "email": "alice@example.com",
  "password": "secret"
}
```

Example response:

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkFsaWNlIiwiZW1haWwiOiJhbGljZUBleGFtcGxlLmNvbSIsImV4cCI6MTYyNDEyMzQwNH0.Q9xL8pTPnlGyvEzJ4Kjh8z9zC8vzGQ8zFJnY5fH9R3k"
}
```

## Endpoints

The API has the following endpoints:

### Menu

- `GET /menu` - Get the list of available pizzas and their prices.
- `GET /menu/{id}` - Get the details of a specific pizza by its ID.

### Orders

- `POST /orders` - Create a new order with one or more pizzas. Requires authentication.
- `GET /orders` - Get the list of orders for the current user. Requires authentication.
- `GET /orders/{id}` - Get the details of a specific order by its ID. Requires authentication and authorization.
- `PATCH /orders/{id}` - Update the status or delivery address of an order. Requires authentication and authorization.
- `DELETE /orders/{id}` - Cancel an order. Requires authentication and authorization.

### Ratings

- `POST /ratings` - Create a new rating for an order. Requires authentication and authorization.
- `GET /ratings` - Get the list of ratings for the current user. Requires authentication.
- `GET /ratings/{id}` - Get the details of a specific rating by its ID. Requires authentication and authorization.
- `PATCH /ratings/{id}` - Update the score or comment of a rating. Requires authentication and authorization.
- `DELETE /ratings/{id}` - Delete a rating. Requires authentication and authorization.

## Data Models

The API uses the following data models:

### Pizza

A pizza is an item on the menu that has a name, a description, a list of toppings, and a price.

Example:

```json
{
  "id": 1,
  "name": "Margherita",
  "description": "A classic pizza with tomato sauce, mozzarella cheese, and basil",
  "toppings": ["tomato sauce", "mozzarella cheese", "basil"],
  "price": 10
}
```

### Order

An order is a request made by a user to buy one or more pizzas. An order has an ID, a user ID, a list of pizzas, a total price, a status, a delivery address, and a timestamp.

Example:

```json
{
  "id": 1,
  "user_id": 1234567890,
  "pizzas": [
    {
      "pizza_id": 1,
      "quantity": 2
    },
    {
      "pizza_id": 3,
      "quantity": 1
    }
  ],
  "total_price": 25,
  "status": "pending",
  "delivery_address": "123 Main Street, Springfield, USA",
  "created_at": "2023-06-19T17:45:04+03:00"
}
```

### Rating

A rating is a feedback given by a user after receiving an order. A rating has an ID, an order ID, a user ID, a score (from 1 to 5), a comment, and a timestamp.

Example:

```json
{
  "id": 1,
  "order_id": 1,
  "user_id": 1234567890,
  "score": 4,
  "comment": "The pizzas were delicious but the delivery was late",
  "created_at": "2023-06-19T18:15:04+03:00"
}
```