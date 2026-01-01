# abigenx demo

Minimal example app that uses generated bindings + event topics to read logs from an Ethereum node.

## Prerequisites

- `abigen` installed
- ABI JSON files for your target contract(s)
- `abigenx` binary built from this repo

## Generate bindings

From `demo/`:

```
make tools
make gen-bindings
```

## Run the demo

From `demo/`:

```
go run . \
  -rpc https://your-rpc.example \
  -address 0xYourContractAddress \
  -from 19000000 \
  -to 19001000
```

This example listens for the `TokenRebased` event on the Lido contract. Swap the import and topics for your own contract as needed.

## Using Outside This Repo

Bindings are generated into `demo/generated/bindings` and imported as `demo/generated/...`.
