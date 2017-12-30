package main

import (
	"blockchain/cli"
	"blockchain/lib"
)

func main() {
	bc := blockchain.NewBlockChain("Aniket")
	defer bc.Close()

	cmd := cli.NewCli(bc)
	cmd.Run()
}
