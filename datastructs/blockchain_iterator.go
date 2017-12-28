package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

// BlockChainIterator represents iterator for getting next block by time
type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bci *BlockChainIterator) Next() *Block {

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
