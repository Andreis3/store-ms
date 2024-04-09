run-app:
	@go run ./cmd/store/store.go
docker-up:
	@docker compose -f docker-compose.yml up -d --build

docker-down:
	@docker compose -f docker-compose.yml down

unit-test:
	@go test ./tests/unit/... --tags=unit -v

unit-tests-cover:
	@go test ./tests/unit/... -coverpkg ./internal/... --tags=unit -v

unit-tests-report:
	mkdir -p "coverage" \
	&& go test ./tests/unit/... -v -coverprofile=coverage/cover.out -coverpkg ./internal/... --tags=unit \
	&& go tool cover -html=coverage/cover.out -o coverage/cover.html \
	&& go tool cover -func=coverage/cover.out -o coverage/cover.functions.html

integration-test:
	@go test ./tests/integration/... --tags=integration -v -count=1

unit-tests-ginkgo:
	@ginkgo -r --race --tags=unit

unit-tests-verbose:
	@ginkgo -r --race --tags=unit

.PHONY: run-app,
		docker-up,
		docker-down,
		unit-test
		unit-tests-cover
		unit-tests-report
		unit-tests-ginkgo
		integration-test