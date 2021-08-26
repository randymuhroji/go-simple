REST_MAIN := "$(CURDIR)/cmd/rest"
BIN_REST := "$(CURDIR)/bin/rest-service"


fetch:
	@go mod tidy
	@go mod download
	@go mod vendor

build-rest:
	@go build -i -v -o $(BIN_REST) $(REST_MAIN)

build-rest-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_REST) $(REST_MAIN)

deploy: clean fetch build-rest

swagger-gen:
	@swag init -g cmd/rest/main.go --output pkg/swagger/docs

clean:
	@rm -rf $(CURDIR)/bin

run-test:
	@go test -v -cover ./... | grep _test.go > output.out