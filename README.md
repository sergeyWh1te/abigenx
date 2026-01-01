# abigenx

Small CLI helper to generate Go bindings from Ethereum ABI JSON files using `abigen`, with optional event topic constants.

## What It Does

- Reads `*.json` ABI files from `-abi-dir`
- Generates Go bindings into `-out-dir` (package = `snake_case` contract name)
- Rewrites methods if ABI uses `ALL_CAPS_WITH_UNDERSCORES`
- Optionally writes `<contract>Events.go` with event names and topic hashes

## Install (Makefile style)

```
BIN_DIR := $(CURDIR)/bin
ABIGEN_BIN := $(BIN_DIR)/abigen
ABIGENX_BIN := $(BIN_DIR)/abigenx

tools:
	@echo "Installing dev tools into ./bin ..."
	GOBIN=$(PWD)/bin go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7
	GOBIN=$(PWD)/bin go install github.com/sergeyWh1te/abigenx/cmd/abigenx@latest
```

Use:

```
$(ABIGENX_BIN) bindings -abi-dir "$(ABI_DIR)" -out-dir "$(BINDINGS_DIR)" -abigen "$(ABIGEN_BIN)" -gen-topics
```

## Usage

```
abigenx bindings -abi-dir ./abi -out-dir ./generated/bindings -abigen ./bin/abigen
abigenx bindings -abi-dir ./abi -out-dir ./generated/bindings -abigen ./bin/abigen -gen-topics
```

Output example:

```
processing /path/to/abi/Lido.json -> /path/to/generated/bindings/mainnet/lido/lido.go
processing /path/to/abi/Lido.json -> /path/to/generated/bindings/mainnet/lido/lidoEvents.go
```

## Demo

The `demo` module shows how to generate bindings and consume event topics in a log reader app.

From `demo/`:

```
make tools
make gen-bindings
go run . -rpc <url> -address <hex> -from <block> -to <block>
```

After `make gen-bindings`, the `demo/generated` folder appears and contains Go bindings for the contract and its events.

Example event processing loop:

```go
for block := fromBlock; block <= toBlock; block++ {
	blockNum := rpc.BlockNumber(block)
	receipts, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHash{BlockNumber: &blockNum})
	if err != nil {
		fail(err.Error() + "\n")
	}

	for _, receipt := range receipts {
		for _, lg := range receipt.Logs {
			if lg.Address != address {
				continue
			}
			if len(lg.Topics) == 0 {
				continue
			}

			foundLogs = true
			switch lg.Topics[0] {
			case lido.Events[lido.TokenRebasedEvent]:
				event, err := filterer.ParseTokenRebased(*lg)
				if err != nil {
					log.Printf("parse token rebased: %v", err)
					continue
				}
				log.Printf(
					"token rebased timeElapsed=%s preTotalShares=%s postTotalShares=%s",
					event.TimeElapsed.String(),
					event.PreTotalShares.String(),
					event.PostTotalShares.String(),
				)
			default:
				log.Printf("unhandled topic=%s", lg.Topics[0].Hex())
			}
		}
	}
}
```

## Notes

- ABI JSON can be either a raw array or an artifact with an `abi` field.
- Event topics are generated using `keccak256` of the event signature.

## Project Layout

- `cmd/abigenx` - CLI entry point
- `pkg/bindings` - binding generation and abigen post-processing
- `pkg/topics` - event topic generation from ABI
- `pkg/naming` - naming helpers
