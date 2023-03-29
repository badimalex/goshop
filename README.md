# Go CRUD Application with PostgreSQL

This is a simple CRUD application built with Go and PostgreSQL. It provides an HTTP API for managing products and users.

## Installation

1. Clone the repository:

```sh
git clone git@github.com:badimalex/goshop.git
```

2. Install dependencies:

```sh
go mod download
```


3. Create a `config.yaml` file with the project configuration (you can use `config.example.yaml` as a template).

4. Start the server:

```sh
go run ./cmd
```

## Usage

The following API endpoints are available:

### User Endpoints

- `POST /register` - register a new user
- `POST /login` - log in a user and get a token

### Product Endpoints

- `POST /products` - create a new product
- `GET /products` - search for products by name (`name` query parameter)
- `GET /products/:id` - get a product by ID
- `PUT /products/:id` - update a product by ID
- `DELETE /products/:id` - delete a product by ID
