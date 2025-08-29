.PHONY: up down
up:
	docker compose up -d
	docker compose logs -t  -f app >> app.log
down:
	docker compose down --rmi local
	rm -f app.log