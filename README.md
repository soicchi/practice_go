# practice_go

# How to run local environment

```sh
docker compose up --build
```

# Command

```sh
# formatter
make go_fmt

# linter
make go_vet

# adjust dependent package
make go_tidy

# Add a cobra command
# ex) make cobra_add name=migrate
make cobra_add name="comamnd name"

# Migration
make migrate

# rollback last migration
make rollback_last

# rollback to specific migration version
# ex) make rollback_to version=202405201334
make rollback_to version="migration_id"
```
