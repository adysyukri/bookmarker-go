css-watch:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --watch=always &
css:
	npx tailwindcss -i public/css/input.css -o public/css/style.css --minify

run: css-watch
	templ generate && air

init-table:
	docker exec -it cockroach1 ./cockroach --host=cockroach1:26357 init --insecure; \
	go run cmd/cockroachinit/main.go

dev: css-watch
	templ generate && air

run-db:
	docker compose -f ./docker/docker-compose-cdb.yml up -d

PHONY: run-api css-watch run css