
# To-Do Social Media API

This project is a RESTful API for managing social media posts, built with Go, Fiber, and GORM. It provides endpoints for creating, reading, updating, and deleting posts, as well as fetching posts within specific date ranges.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features

- CRUD operations for social media posts
- Fetch posts by date range
- Get posts due within a specified number of days
- API rate limiting
- Swagger documentation
- PostgreSQL database integration

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.16 or later)
- PostgreSQL
- Git

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/rrraf1/to-do_social_media.git
   cd to-do_social_media
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your environment variables by copying the `.env.example` file to `.env` and filling in your database details:

   ```bash
   cp .env.example .env
   ```

   Then edit the `.env` file with your PostgreSQL credentials.

4. Create the database:

   ```bash
   createdb db_name
   ```

5. Run the application:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:3000` by default.

## Usage

Once the server is running, you can interact with the API using HTTP requests. Here are some example endpoints:

- GET `/posts`: Retrieve all posts
- POST `/posts`: Create a new post
- GET `/posts/{targetDate}`: Get posts due within the next `targetDate` days
- PUT `/posts/{id}`: Update a post
- DELETE `/posts/{id}`: Delete a post

For a complete list of endpoints and their usage, refer to the [API Documentation](#api-documentation) section.

## API Documentation

This project uses Swagger for API documentation. Once the server is running, you can access the Swagger UI at:

```
http://localhost:3000/swagger/
```

This interactive documentation provides detailed information about each endpoint, including:

- Available methods
- Request parameters
- Request body schemas
- Response schemas
- Example requests and responses

You can also test the API directly from the Swagger UI.

## Project Structure

```
.
├── connection/
│   └── connection.go    # Database connection setup
├── controller/
│   └── posts.go         # Request handlers for posts
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml     # Swagger documentation files
├── model/
│   ├── post.go          # Post model definition
│   └── swagger.go       # Swagger response models
├── routers/
│   └── routes.go        # API route definitions
├── .env                 # Environment variables (git-ignored)
├── .gitignore
├── go.mod
├── go.sum
├── main.go              # Application entry point
└── README.md            # This file
```

## Environment Variables

The following environment variables should be set in your `.env` file:

- `DB_HOST`: PostgreSQL host (e.g., "localhost")
- `DB_USER`: PostgreSQL user
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `PORT`: (Optional) Port for the server to listen on (default: 3000)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open-source and available under the.

```
