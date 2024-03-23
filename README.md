# loop-app
A simple pulse-survey application.

## Getting Started
### Prerequisites
- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)

### Setup
1. Clone the repository.
```bash
git clone github.com/jbactad/loop.git
cd loop
```

2. Install tools.
```bash
make setup
```

3. Start the dependencies, i.e. database.
```bash
make dep-up
```

4. Run the migrations.
```bash
make db-migrate
make db-seed # Optional
```

5. Start the application.
```bash
make run
```

Go to [http://localhost:8080](http://localhost:8080) to view GraphQL Playground.

## Project Structure
```text
.
├── application
│   ├── commands
│   ├── ports
│   └── queries
├── cmd
│   ├── db
│   └── graphql
├── db
│   ├── migrations
│   └── seeds
├── domain
├── graph
│   ├── generated
│   ├── models
│   ├── resolvers
│   └── schemas
│   └── tests
└── infrastructure
    └── repositories
```

- `application`: Contains the application layer. This layer contains the use cases, which are the application-specific business rules.
- `cmd`: Contains different commands that can be run from the command line.
- `db`: Contains the database migrations and seeds.
- `domain`: Contains the domain layer. This layer contains the domain model and the domain services.
- `graph`: Contains the GraphQL layer. This layer contains the generated code, the models, the resolvers, the schemas, and the tests.
- `infrastructure`: Contains the infrastructure layer. This layer contains the implementation of the ports defined in the application layer.

