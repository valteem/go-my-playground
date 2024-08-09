.PHONY: debug
debug: cmd/main.go server/server.go
	@go build \
		-gcflags "all=-N -l" \
		-o echo-crud \
		./cmd
		

.PHONY: clean
clean:
	@rm echo-crud