go_fmt:
	docker compose run --rm api go fmt ./...

go_vet:
	docker compose run --rm api go vet ./...

go_tidy:
	docker compose run --rm api go mod tidy

cobra_add:
	docker compose run --rm api cobra-cli add ${name}

migrate:
	docker compose run --rm api go run main.go migrate

rollback_last:
	docker compose run --rm api go run main.go rollbackLast

rollback_to:
	docker compose run --rm api go run main.go rollbackTo ${version}
