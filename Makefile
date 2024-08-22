css-watch:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch=always &
css:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --minify

run: css-watch
	templ generate && air

init-table:
	go run cmd/cockroachdb/main.go
	
PHONY: run-api css-watch run css