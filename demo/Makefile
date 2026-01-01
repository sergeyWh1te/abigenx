BIN_DIR := $(CURDIR)/bin
ABIGEN ?= $(BIN_DIR)/abigen
ABIGENX ?= $(BIN_DIR)/abigenx

ABI_DIR := $(CURDIR)/abi
BINDINGS_DIR := $(CURDIR)/generated/bindings

.PHONY: tools
tools:
	@echo "Installing tools into ./bin ..."
	GOBIN=$(PWD)/bin go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7
	GOBIN=$(PWD)/bin go install github.com/sergeyWh1te/abigenx/cmd/abigenx@latest

.PHONY: gen-bindings
gen-bindings:
	$(ABIGENX) bindings -abi-dir "$(ABI_DIR)" -out-dir "$(BINDINGS_DIR)" -abigen "$(ABIGEN)" -gen-topics

.PHONY: run
run:
	go run . -rpc "$(RPC)" -address "$(ADDRESS)" -from "$(FROM)" -to "$(TO)"
