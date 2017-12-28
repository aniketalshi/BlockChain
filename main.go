package main

import (
	"blockchain/datastructs"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("BTC 1")
	bc.AddBlock("BTC 2")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev Hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Println()
	}
}
