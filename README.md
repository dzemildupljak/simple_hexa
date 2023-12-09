# My Hexagonal Go App

This is a simple example of a Hexagonal Architecture in a Go application.

## Project Structure

The project follows a Hexagonal Architecture, also known as Ports and Adapters. Here's a brief explanation of the folder structure:

- **cmd:** This folder contains the main entry point of the application. The `main.go` file initializes the application and starts the execution.

- **internal:** The internal code of the application is organized into `app` and `infrastructure`.

  - **app:**

    - **application:** Contains interfaces and implementations for application services.
    - **domain:** Defines business domain models and repository interfaces.
    - **ports:**
      - **inbound:** Inbound ports define interfaces for communication with the external world (e.g., API handlers).
      - **outbound:** Outbound ports define interfaces for external dependencies (e.g., database interfaces).

  - **infrastructure:** Implements infrastructure details, such as database access or external service connections.

- **pkg:** This directory is for reusable packages that can be shared across different parts of the application.

- **scripts:** Includes build and deployment scripts.

- **config:** Configuration files for the application.

- **tests:** Test files for the application.

- **go.mod and go.sum:** These files manage the project's dependencies.
