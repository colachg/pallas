# Pallas
This project is a practice of learning golang and graphql.Soon I want to create a tutorial of the combo I used in this.

## Tools
- gqlgen
- migrate
- gorm
- sqlmock
- echo
- go-sqlmock
 
## Get start
1. `mkdir pallas`
2. `go mod init`
3. create `schema.graphql` and add custom types, query and mutation.
4. `go run github.com/99designs/gqlgen init` to generate the skeleton.
5. modify gqlgen.yml to set the project structure as you like.
6. `go run github.com/99designs/gqlgen -v` to regenerate files.
7. Create database model and add it to `gqlgen.yml` and regenerate.

## Database tools
1. `migrate create -ext sql -dir mysql/migrations create_projects`
2. `migrate -source file://path/to/migrations -database mysql://user:password@tcp(host:port)/dbname up`
3. `migrate -source file://path/to/migrations -database mysql://user:password@tcp(host:port)/dbname down`

## Todo
- [x] Add `echo` web framework
- [x] Refactor test suites
- [ ] Add auth middleware
- [ ] Update README.md
- [ ] Dataloaden
- [ ] Write tutorial