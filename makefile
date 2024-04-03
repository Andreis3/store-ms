run-app:
	@go run ./cmd/store/store.go
docker-up:
	@docker compose -f ./docker/docker-compose.yml up -d --build

docker-down:
	@docker compose -f ./docker/docker-compose.yml down

unit-test:
	@go test ./tests/unit/... --tags=unit -v

unit-tests-cover:
	@go test ./tests/unit/... -coverpkg ./... --tags=unit

unit-tests-report:
	mkdir -p "coverage" \
	&& go test ./tests/... -v -coverprofile=coverage/cover.out -coverpkg ./... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

unit-tests-ginkgo:
	@ginkgo -r --race --tags=unit

.PHONY: run-app,
		docker-up,
		docker-down,
		unit-test
		unit-tests-cover
		unit-tests-cover
		unit-tests-report
		unit-tests-ginkgo