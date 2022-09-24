APP_NAME = twitter-translate

build:
	@go build -o ${APP_NAME} ./cmd/${APP_NAME}

test:
	@go test ./...

.PHONY: integration
integration:
	@go test -tags=integration ./...
