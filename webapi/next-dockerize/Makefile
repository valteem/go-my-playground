.PHONY: clean
clean:
	@rm -rf  ./bin/

.PHONY: docker-build
docker-build:
	@docker build -t next-app .

.PHONY: docker-run
docker-run:
	@docker run -p 3001:3001 --name next-app -d next-app

.PHONY: docker-build-dev
docker-build-dev:
	@docker build -f Dockerfile.dev -t next-app-dev .

.PHONY: docker-run-dev
docker-run-dev:
	@docker run -p 3001:3001 -v `pwd`/:/app/ --name next-app-dev -dit next-app-dev

.PHONY: compose-down
compose-down:
	@docker-compose down --remove-orphans --rmi local

.PHONY: compose-up
compose-up:
#	@docker-compose up --build -d && docker-compose logs -f
	@docker-compose up --build -d 