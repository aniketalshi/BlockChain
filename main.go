package main

import (
	"blockchain/datastructs"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("BTC 1")
	bc.AddBlock("BTC 2")
}
