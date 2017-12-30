package blockchain

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

const dbFile = "blocks.db"
const blocksBucket = "Blocks"
const genesisCoinBaseData = "genesis Txn"

// BlockChain defines link of blocks
type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

// Close exports close functionality of db
func (b *BlockChain) Close() {
	b.db.Close()
}

// AddBlock appends the block to end of blockchain
func (b *BlockChain) AddBlock(tx *Transaction) {
	// construct new block and prev hash will be current tip of db
	block := NewBlock([]*Transaction{tx}, b.tip)

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

//NewGenesisBlock creates first BlockChain block
func NewGenesisBlock(tx *Transaction) *Block {
	return NewBlock([]*Transaction{tx}, []byte{})
}

//NewBlockChain constructs a new block chain
func NewBlockChain(address string) *BlockChain {
	//return &BlockChain{[]*Block{GenerateGenesisBlock()}}

	var tip []byte

	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		// if we dont already have bucket with blocks
		if b == nil {
			fmt.Println("Creating New Blockchain.")
			newTx := NewCoinBase(address, genesisCoinBaseData)
			bl := NewGenesisBlock(newTx)

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

// GetIterator returns pointer to iterate upon the blockchain
func (b *BlockChain) GetIterator() *Iterator {
	return &Iterator{b.tip, b.db}
}

// Print is utility to print info of blocks
func (b *BlockChain) Print() {

	iter := b.GetIterator()

	for {
		b := iter.Next()
		if b == nil || len(b.PrevBlockHash) == 0 {
			break
		}
		b.Print()
	}
}
