package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sergeyWh1te/abigenx/pkg/bindings"
)

const usage = "usage: abigenx bindings -abi-dir <dir> -out-dir <dir> -abigen <path> [-gen-topics]"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fail(usage)
	}

	if !strings.HasPrefix(args[0], "-") {
		if args[0] != "bindings" {
			fail("unsupported command: " + args[0])
		}
		args = args[1:]
	}

	flags := flag.NewFlagSet("abigenx", flag.ExitOnError)
	var abiDir, outDir, abigenPath string
	var genTopics bool
	flags.StringVar(&abiDir, "abi-dir", "", "directory with ABI json files")
	flags.StringVar(&outDir, "out-dir", "", "output directory for bindings")
	flags.StringVar(&abigenPath, "abigen", "abigen", "path to abigen binary")
	flags.BoolVar(&genTopics, "gen-topics", false, "generate event topics files")
	flags.Parse(args)

	if abiDir == "" || outDir == "" {
		fail(usage)
	}

	if err := bindings.Generate(abiDir, outDir, abigenPath, genTopics); err != nil {
		fail(err.Error())
	}
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
