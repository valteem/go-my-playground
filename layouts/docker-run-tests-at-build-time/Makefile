.PHONY: test-stage-up
test-stage-up:
	@docker compose -f test-stage.yml up --build -d 

.PHONY: test-stage-down
test-stage-down:
	@docker compose -f test-stage.yml down --rmi local

.PHONY: test-cmd-up
test-cmd-up:
	@docker compose -f test-cmd.yml up --build

.PHONY: test-cmd-down
test-cmd-down:
	@docker compose -f test-cmd.yml down --rmi local