# Go Notifier

Go Notifier is a notification service written in Go that leverages a bot to send messages based on various triggers or
events. The service is designed to be easily configurable and extendable, making it suitable for various notification
use cases.

Example: https://t.me/go_versions

## Features

- **Bot Integration**: Utilizes a bot to send notifications.
- **Customizable Templates**: Supports customizable message templates using Markdown.
- **Environment Configuration**: Configurable via `.env` files for easy setup.
- **Docker Support**: Includes Dockerfile and Docker Compose for containerized deployments.

## Project Structure

- **`cmd/notifier/main.go`**: The entry point of the application.
- **`internal/`**: Contains core logic for the bot and parser functionalities.
    - `bot.go`: Contains the logic for interacting with the bot.
    - `parser.go`: Contains parsing logic for handling data.
    - `parser_test.go`: Unit tests for the parser.
- **`templates/`**: Directory for storing message templates.
- **`.env`**: File for environment-specific configuration.
- **`Dockerfile`**: Dockerfile for building the Go application.
- **`docker-compose.yml`**: Docker Compose file for setting up and running the application in a containerized
  environment.

## Getting Started

### Prerequisites

- Go 1.23 or later
- Docker (optional, for containerized deployment)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-notifier.git
   cd go-notifier
   ```

2. Set up environment variables:

   Copy the `.env.example` to `.env` and fill in the required variables.

   ```bash
   cp .env.example .env
   ```

3. Build and run the application:

   ```bash
   go build -o notifier ./cmd/notifier
   ./notifier
   ```

### Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t go-notifier .
   ```

2. Run the container:

   ```bash
   docker-compose up
   ```

## Usage

- Configure your message templates in the `templates` folder.
- Adjust environment variables in the `.env` file to suit your needs.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
