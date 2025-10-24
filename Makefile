build:
	@go test
	@go clean
	@go vet
	@go build .
	@scss frontend/assets/scss/*.scss frontend/assets/css/style.css
