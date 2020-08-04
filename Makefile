.PHONY: all build clean install uninstall fmt simplify check run test

install:
	@go install ./bin/run.go

run: install
	@go run ./bin/run.go

test:
	@go test -v ./test/...