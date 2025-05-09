.PHONY: clean
clean:
	@rm -rf  ./bin/

.PHONY: docker-build
docker-build:
	@docker build -t next-app .

.PHONY: docker-run
docker-run:
	@docker run -p 3001:3001 --name next-app -d next-app