package main

import (
	"blockchain/lib"
)

func main() {
	bc := blockchain.NewBlockChain()
	defer bc.Close()

	cli := blockchain.NewCli(bc)
	cli.Run()
}
