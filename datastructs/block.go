package blockchain

import (
	"bytes"
	"encoding/gob"
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

	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
