run-app:
	@go run ./cmd/store/store.go
docker-up:
	@docker compose -f ./docker/docker-compose.yml up -d --build

docker-down:
	@docker compose -f ./docker/docker-compose.yml down

unit-test:
	@go test ./tests/unit/... --tags=unit -v


.PHONY: run-app,
		docker-up,
		docker-down,
		unit-test