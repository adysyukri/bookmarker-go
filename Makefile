run-api:
	cd cmd/api && go run *.go

css-watch:
	npx tailwindcss -i web/assets/css/input.css -o web/assets/css/style.css --watch &

histo-dev:
	cd cmd/duohistopath && air
PHONY: run-api css-watch histo-dev