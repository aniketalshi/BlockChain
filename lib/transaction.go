package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const subsidy = 5

// Transaction denotes single txn in blockchain
type Transaction struct {
	ID  []byte
	In  []TxInput
	Out []TxOutput
}

// TxInput struct defines input of given txn
type TxInput struct {
	ID        []byte // id of Transaction
	Out       int    // index of output in Transaction
	ScriptSig string // used for unlocking input's ScriptPubKey
}

// TxOutput defines output of given txn
type TxOutput struct {
	Value        int64  // value for mining
	ScriptPubKey string // key for unlocking this output
}

// Serialize encodes Transaction into bytes
func (t *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(t); err != nil {
		log.Fatal("serialize :", err)
	}
	return result.Bytes()
}

// SetID sets the Transaction id to hash of itself
func (t *Transaction) SetID() {
	hash := sha256.Sum256(t.Serialize())
	t.ID = hash[:]
}

// NewCoinBase generates first Transaction mining Genesis block
func NewCoinBase(to, data string) *Transaction {
	inp := TxInput{[]byte{}, -1, data}
	oup := TxOutput{subsidy, to}

	tx := Transaction{nil, []TxInput{inp}, []TxOutput{oup}}
	tx.SetID()
	return &tx
}
