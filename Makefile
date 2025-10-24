build:
	@templ generate
	@go test
	@go clean
	@go vet
	@go build .
	@scss frontend/assets/scss/main.scss frontend/assets/css/style.css
