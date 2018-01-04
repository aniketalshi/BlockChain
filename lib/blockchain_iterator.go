package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
)

// Iterator represents iterator for getting next block by time
type Iterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next gives us next block in the chain
func (bci *Iterator) Next() *Block {

	var block *Block
	err := bci.db.View(func(tx *bolt.Tx) error {
		bckt := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bckt.Get(bci.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Fatal("BlockChainIterator Next:", err)
	}

	bci.currentHash = block.PrevBlockHash
	return block
}
