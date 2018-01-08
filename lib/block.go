package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Block represents single block of data in chain storing hashesh
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock is constructor taking in data and prev hash combining them into a new struct.
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {

	fmt.Printf("Mining New Block...\n")
	block := &Block{time.Now().Unix(), txs, prevBlockHash, []byte{}, 0}

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

	fmt.Printf("\n\nNew Block :: Timestamp : %d, Hash :%x\n", b.Timestamp, b.Hash)
	
	fmt.Println("--- Printing all transactions ---");
	for _, txn := range b.Transactions {
				
		for _, inTxn := range txn.In {
			inTxn.Print()
		}

		for _, outTxn := range txn.Out {
			outTxn.Print()
		}
	}
}

// HashTransactions concatenates all hashes of txn and creates unique hash to identify them
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]

}
