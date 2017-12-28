package blockchain

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

const dbFile = "blocks.db"
const blocksBucket = "Blocks"

// BlockChain defines link of blocks
type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock appends the block to end of blockchain
func (b *BlockChain) AddBlock(data string) {
	// construct new block and prev hash will be current tip of db
	block := NewBlock(data, b.tip)

	err := b.db.Update(func(tx *bolt.Tx) error {
		bckt := tx.Bucket([]byte(blocksBucket))
		if err := bckt.Put(block.Hash, block.Serialize()); err != nil {
			return err
		}
		if err := bckt.Put([]byte("l"), block.Hash); err != nil {
			return err
		}
		b.tip = block.Hash
		return nil
	})

	if err != nil {
		log.Fatal("AddBlock :", err)
	}
}

//GenerateGenesisBlock creates first BlockChain block
func GenerateGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//NewBlockChain constructs a new block chain
func NewBlockChain() *BlockChain {
	//return &BlockChain{[]*Block{GenerateGenesisBlock()}}

	var tip []byte

	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		// if we already have bucket with blocks
		if b == nil {
			bl := GenerateGenesisBlock()

			bckt, _ := tx.CreateBucket([]byte(blocksBucket))
			if newerr := bckt.Put(bl.Hash, bl.Serialize()); err != nil {
				return newerr
			}
			if newerr := bckt.Put([]byte("l"), bl.Hash); err != nil {
				return newerr
			}
			tip = bl.Hash
		} else {
			// set tip to last existing hash
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Fatal("NewBlockChain :", err)
	}
	return &BlockChain{tip, db}
}
