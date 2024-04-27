css-watch:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch
css:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --minify

run: css-watch
	air
PHONY: run-api css-watch run css