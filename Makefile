fmt:
	@gofmt -w -s . && goimports -w . && go vet ./... && go mod tidy
test:
	@mkdir -p out/ && go test ./... -count=1 -coverprofile=out/coverage.out
test-cover: test
	@go tool cover -html=out/coverage.out