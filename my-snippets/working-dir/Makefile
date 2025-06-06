.PHONY: ltest
ltest:
	go test -v ./...

.PHONY: lrun
lrun:
	go run ./level1/level2/level3/cmd/main.go
	cd level1 && go run ./level2/level3/cmd/main.go
	cd level1/level2 && go run ./level3/cmd/main.go
	cd level1/level2/level3 && go run ./cmd/main.go
	cd level1/level2/level3/cmd && go run ./main.go

.PHONY: lbin
lbin:
	@echo "creating /bin directory ..."
	@mkdir -p bin
	@echo "\nbuilding executable and placing it to /bin ..."
	@go build -o ./bin/a.out ./level1/level2/level3/cmd/main.go
	@echo "\nrunning executable from root project directory ..."
	./bin/a.out
	@echo "\nrunning executable from /bin directory ..."
	cd bin && ./a.out
	@echo "\nrunning executable from /sibling directory ..."
	cd sibling && ./../bin/a.out

,PHONY: clean
clean:
	@rm -rf bin