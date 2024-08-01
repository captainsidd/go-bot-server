run:
	make -j 2 local-server local-test
.PHONY: local

local-server:
	@go run main.go
.PHONY: local-server

local-test:
	@sh scripts/test.sh
.PHONY: local-test
