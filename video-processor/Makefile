BUILD_TARGET := build/video-processing-api

build:
	go build -o $(BUILD_TARGET) cmd/video-processor/main.go
	
test:
	go test ./...
	
run:
	./$(BUILD_TARGET)


.PHONY: build test run
