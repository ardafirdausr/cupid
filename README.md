# Cupid

Cupid is Tinder-like backend service for matchmaking. It is a RESTful API that allows users to create accounts, view other users, and match with other users.

## Run the project

1. Clone this repository: `git clone git@github.com:ardafirdausr/cupid.git`
2. Copy `.env.example` to `.env`
3. Fill the environment variables in the `.env` file
4. Install go dependencies

   ```bash
   go get
   ```

5. Start the app

    ```bash
    go run main.go
    ```  

## Development

### Generate code

This project using some generated code. To start development, you have to install the following tools:

1. [Mockgen](https://github.com/golang/mock)  
   This tool is used to generate mock for the interface. To install this tool, run the following command:

   ```bash
   go install github.com/golang/mock/mockgen@v1.6.0
   ```

2. [Wire](https://github.com/google/wire)  
   This tool is used to generate dependency injection. To install this tool, run the following command:

   ```bash
   go install github.com/google/wire/cmd/wire@latest
   ```

After installing the tools, you can run the following command to generate the mock, dependency injection, and swagger
documentation:

```bash
  go generate ./...
```

### Git Branching Strategy

This repository uses trunk-based development branching strategy.

- All feature branches are branched off from `main` and merged back into `main`.
- All feature branches are short-lived and should be merged back into `main` as soon as possible.
- All feature branches are prefixed with `feat/`.
- Main branch will always be deployable as staging.
- Production environtment will be deployable on release branch with tagging version.

For more information on trunk-based development, see [trunkbaseddevelopment.com](https://trunkbaseddevelopment.com/).

## Software Architecture

This project uses monolithic architecture for the mvp version. The reason for choosing a monolithic architecture is because it is easier to develop and deploy. The app is also small enough that a monolithic architecture is sufficient.  

For the software architecture, we will be using clean architecture. Clean architecture is a software design philosophy that separates the software into layers. Each layer has a specific responsibility and interacts with other layers in a specific way. The layers are:

- Entities
- Use Cases
- Interface Adapters
- App (the most outer layer)

For more information on clean architecture, see [cleanarchitecture.io](https://www.cleancoders.com/episode/clean-code-episode-42/show).

### Directory Structure

```plaintext
├── app                     # The way to start the app (http, cli, grpc, cron, etc)
│   └── http                # The http server to serve the app
│       ├── handler         # The http handler for the app
│       │   └── response    # The response for the http handler
│       └── middleware      # The middleware for the http server
├── docs                    # The documentation for the project
└── internal                # The internal package that contains the core of the project
    ├── dto                 # The data transfer object used to transfer data between layers
    ├── entity              # The entity for the project (domain)
    │   └── errs            # The error entity for the project
    ├── helper              # The helper codes
    ├── mock                # The mock for core project (usecase, repository, etc)
    ├── pkg                 # Libraray wrapper or custom library based on usecase
    │   ├── echo            # The echo library wrapper
    │   ├── logger          # The logger library wrapper
    │   ├── mongo           # The mongo library wrapper
    │   ├── validator       # Validator library wrapper
    │   └── [other]         # Other library
    ├── repository          # The repository implementation (data access layer, it can be database, cache, service, etc)
    │   ├── mongo           # The mongo repository implementation
    │   │   └── seed        # The seed data for mongo repository
    │   └── [other]         # Other repository implementation
    ├── service             # The service implementation 
    ├── repository.go       # The repository interface for the project
    └── service.go          # The service interface for the project
```
