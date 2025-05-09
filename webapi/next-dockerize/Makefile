.PHONY: clean
clean:
	@rm -rf  ./bin/

.PHONY: docker-build
docker-build:
	@docker build -t next-app .

.PHONY: docker-run
docker-run:
	@docker run -p 3001:3001 --name next-app -d next-app

.PHONY: compose-down
compose-down:
	@docker-compose down --remove-orphans --rmi local

.PHONY: compose-up
compose-up:
#	@docker-compose up --build -d && docker-compose logs -f
	@docker-compose up --build -d 