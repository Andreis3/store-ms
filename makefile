docker-up:
	@docker compose -f ./docker/docker-compose.yml up -d --build

docker-down:
	@docker compose -f ./docker/docker-compose.yml down

unit-test:
	@go test ./tests/unit/... --tags=unit -v


.PHONY: docker-up,
		docker-down,
		unit-test