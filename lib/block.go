package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Block represents single block of data in chain storing hashesh
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock is constructor taking in data and prev hash combining them into a new struct.
func NewBlock(data string, prevBlockHash []byte) *Block {

	fmt.Printf("Mining New Block...\n")
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	block.Print()
	fmt.Println("Success")

	return block
}

//Serialize converts block struct into bytes
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(b); err != nil {
		log.Fatal("serialize :", err)
	}
	return result.Bytes()
}

// Print is utility function to print block
func (b *Block) Print() {
	fmt.Printf("Timestamp : %d, Data : %s, Hash :%x\n", b.Timestamp, string(b.Data[:]), b.Hash)
}
