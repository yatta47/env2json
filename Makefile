NAME := env2json
VERSION := 0.0.1
BIN_DIR := bin

.PHONY: run
run:
	@go run main.go

.PHONY: build
build:
	@go build -o $(BIN_DIR)/$(NAME) main.go

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)
