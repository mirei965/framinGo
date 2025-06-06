## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool framingo and copies it to myapp
build_cli:
	@go build -o ../myapp/framinGo ./cmd/cli

## build_cli: builds the command line tool framingo and copies it to myapp for Windows
build_cli_windows:
	@go build -o ../myapp/framinGo.exe ./cmd/cli

## build: builds the command line tool to dist directoory
build:
	@go build -o ./dist/framinGo ./cmd/cli