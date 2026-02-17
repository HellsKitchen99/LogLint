info:
	cloc .
run-test:
	go test -cover 
coverage:
	go test -coverprofile=coverage.out 
	go tool cover -html=coverage.out
go-version:
	go version
golangci-lint-version:
	golangci-lint --version
build-plugin:
	go build -o loglint ./cmd/loglint
run-linter:
	go vet -vettool=./loglint ./example