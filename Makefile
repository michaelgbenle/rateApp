run: |
	gofmt -w .
	go run ./cmd/main.go

mock:
	mockgen -source=internal/ports/repository.go -destination=internal/repository/mocks/db_mock.go -package=mocks

tests:
	mockgen -source=internal/ports/repository.go -destination=internal/repository/mocks/db_mock.go -package=mocks
	go test ./internal/api/test/... -v