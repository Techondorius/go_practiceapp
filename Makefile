d:
	docker compose -f compose.dev.yml up -d
dd:
	docker compose -f compose.dev.yml down
db:
	docker compose -f compose.dev.yml build
