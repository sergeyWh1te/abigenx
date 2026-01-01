# abigenx

Small CLI helper to generate Go bindings from Ethereum ABI JSON files using `abigen`, with optional event topic constants.

## What and Why

`abigenx` is a thin wrapper on top of `abigen`. It reads ABI JSON files from a directory, generates Go bindings into a target directory, and (optionally) generates a contract event mapping with topic hashes. Under the hood it calls `abigen` for each ABI.

It is designed to be a tiny, independent tool you can vendor or publish as its own module.

## What It Does

For each `*.json` ABI file under `-abi-dir`:

- Creates a package directory under `-out-dir` (uses `snake_case` of the contract name)
- Generates a Go binding with `abigen`
- Rewrites method names if ABI uses ALL_CAPS_WITH_UNDERSCORES
- Optionally writes `<contract>Events.go` with event names and topic hashes

## Usage

Basic:

```
# generate bindings only
abigenx bindings \
  -abi-dir ./abi \
  -out-dir ./generated/bindings \
  -abigen ./bin/abigen
```

With event topics:

```
abigenx bindings \
  -abi-dir ./abi \
  -out-dir ./generated/bindings \
  -abigen ./bin/abigen \
  -gen-topics
```

Output example:

```
processing /path/to/abi/Lido.json -> /path/to/generated/bindings/mainnet/lido/lido.go
processing /path/to/abi/Lido.json -> /path/to/generated/bindings/mainnet/lido/lidoEvents.go
```

## Install (Makefile style)

You can install `abigenx` alongside other tooling with a Makefile target similar to your setup.

Example:

```
BIN_DIR := $(CURDIR)/bin
ABIGEN_BIN := $(BIN_DIR)/abigen
ABIGENX_BIN := $(BIN_DIR)/abigenx

tools:
	@echo "Installing dev tools into ./bin ..."
	GOBIN=$(PWD)/bin go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7
	GOBIN=$(PWD)/bin go install github.com/sergeyWh1te/abigenx/cmd/abigenx@latest
```

Then use it:

```
$(ABIGENX_BIN) bindings -abi-dir "$(ABI_DIR)" -out-dir "$(BINDINGS_DIR)" -abigen "$(ABIGEN_BIN)" -gen-topics
```

## How to Use in Monitoring / Ethereum Projects

1. Collect ABI JSON files for the contracts you want to monitor or call.
2. Run `abigenx` to generate Go bindings and event topic constants.
3. Import the generated packages in your services to decode logs and call contract methods.

Demo app: `demo`

## Notes

- ABI JSON can be either a raw array or an artifact with an `abi` field.
- Event topics are generated using `keccak256` of the event signature.

## Project Layout

- `cmd/abigenx` - CLI entry point (flag parsing, dispatch)
- `pkg/bindings` - binding generation and abigen post-processing
- `pkg/topics` - event topic generation from ABI
- `pkg/naming` - naming helpers (snake/camel case)
