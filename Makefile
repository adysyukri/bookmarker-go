css-watch:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch=always &
css:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --minify

run: css-watch
	templ generate && air

init-table:
	go run cmd/cockroachinit/main.go

dev: 
	air

PHONY: run-api css-watch run css