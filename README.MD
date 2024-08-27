## Main Commands


### Run: install
```
    make install
```
- _Install all project's dependecies_


### Build
```
    make build
```
- _Build the application_
- _OBS:_ The build file will be available on /bin folder


### Run App - DEV
```
    air
```
- _Run the application with watch mode_


### Migrate DB
```
    make db-migrate-up
```
- _Run migrations_


### Rollback Migration
```
    make db-migrate-down
```
- _Rollback last migration_


### Reset Migrations
```
    make db-migrate-reset
```
- _Reset migrations_


### Create Migration
```
    make db-create-migration MIGRATION_NAME=<NAME>
```
- _Create a new migration_


### Seed DB
```
    make db-seed
```
- _Seed DB_


### Create Seed
```
    make db-create-seed SEED_NAME=<NAME>
```
- _Create a new seed_


### Compile REST API Documentation
```
    make swag-compile
```
- _Compile the API's documentation using Swagger_
- _OBS: The docs is available on /docs/_



#### External Packages Explanation

- Goose: To handle migrations and seeds
- Swaggo: To documentate API routes
- Gorilla Mux: To implement a request router and dispatcher for matching incoming requests to their respective handler 