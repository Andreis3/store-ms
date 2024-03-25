docker-up:
	@docker compose -f ./docker/docker-compose.yml up -d --build

docker-down:
	@docker compose -f ./docker/docker-compose.yml down


.PHONY: docker-up,
		docker-down