build:
	@go build -o go_api cmd/main.go
test: 
	@go test -v ./...
run: build
	@./go_api