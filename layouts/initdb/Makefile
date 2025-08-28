.PHONY: up down
up:
	docker compose up -d
down:
	docker compose down --rmi local
	docker volume rm initdb_pg_initdb_data
