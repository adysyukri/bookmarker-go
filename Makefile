css-watch:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch &

run: css-watch
	air
PHONY: run-api css-watch run