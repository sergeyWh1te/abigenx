package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"demo/generated/bindings/mainnet/lido"
)

func main() {
	var rpcURL string
	var addressHex string
	var fromBlock uint64
	var toBlock uint64

	flag.StringVar(&rpcURL, "rpc", "", "RPC URL")
	flag.StringVar(&addressHex, "address", "", "contract address")
	flag.Uint64Var(&fromBlock, "from", 0, "start block")
	flag.Uint64Var(&toBlock, "to", 0, "end block (defaults to from)")
	flag.Parse()

	if rpcURL == "" || addressHex == "" || fromBlock == 0 {
		fail("usage: demo -rpc <url> -address <hex> -from <block> [-to <block>]\n")
	}
	if toBlock == 0 {
		toBlock = fromBlock
	}
	if toBlock < fromBlock {
		fail("-to must be >= -from\n")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		fail(err.Error() + "\n")
	}
	defer client.Close()

	address := common.HexToAddress(addressHex)
	filterer, err := lido.NewLidoFilterer(address, client)
	if err != nil {
		fail(err.Error() + "\n")
	}

	ctx := context.Background()
	foundLogs := false

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

	if !foundLogs {
		log.Println("no logs found")
	}
}

func fail(msg string) {
	fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}
