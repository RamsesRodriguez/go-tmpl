run:
	@npm install
	@npx tailwindcss -i ./assets/base.css -o ./assets/css/styles.css
	@templ generate
	@go run cmd/main.go