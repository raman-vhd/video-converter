BUILD_TARGET := build/video-converter-api

build:
	go build -o $(BUILD_TARGET) cmd/server/main.go
	
test:
	go test ./...
	
run:
	./$(BUILD_TARGET)


.PHONY: build test run
