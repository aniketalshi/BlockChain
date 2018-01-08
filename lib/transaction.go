package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
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
	ID        []byte // id of output Transaction
	Out       int    // index of output in Transaction
	ScriptSig string // used for unlocking output's ScriptPubKey
}

// TxOutput defines output of given txn
type TxOutput struct {
	Value        int64  // value for mining
	ScriptPubKey string // key for unlocking this output
}

// CheckInputUnlock verifies if input can unlock given ScriptPubKey
func (txIn *TxInput) CheckInputUnlock(key string) bool {
	return txIn.ScriptSig == key
}

// CheckOutputUnlock verifies if output can be unlocked with given key
func (txOup *TxOutput) CheckOutputUnlock(key string) bool {
	return txOup.ScriptPubKey == key
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

	fmt.Printf("The coinbase has been set to %s\n", to)
	return &tx
}

// IsCoinbase verifies if this is first Txn
func (t *Transaction) IsCoinbase() bool {
	if len(t.In) != 1 {
		return false
	}

	txn := t.In[0]
	if txn.Out == -1 && txn.ScriptSig == genesisCoinBaseData {
		return true
	}
	return false
}

// NewUserTransaction creates new transaction sending coins from - to
func NewUserTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	// find all outputs we can spend
	validAmt, validOutputs := bc.FindSpendableOutputs(from, amount)
	if validAmt < amount {
		log.Panic("Address : ", from, " amount requested :", amount, " balance in acc : ", validAmt)
	}

	// go over the map of txnId and outputs
	for txID, outputs := range validOutputs {
		id, _ := hex.DecodeString(txID)

		// go over all the outputs in this txnID and construct inputs for our txn
		for _, out := range outputs {
			inputs = append(inputs, TxInput{id, out, from})
		}
	}

	// build a list of outpus
	outputs = append(outputs, TxOutput{int64(amount), to})

	// give change back to owner
	if validAmt > amount {
		outputs = append(outputs, TxOutput{int64(validAmt - amount), from})
	}

	finalTxn := Transaction{nil, inputs, outputs}
	finalTxn.SetID()

	return &finalTxn
}

func (inTxn *TxInput) Print() {
	fmt.Printf("INPUT TXN ::\n ID : %x From : %s\n", inTxn.ID, inTxn.ScriptSig)
}

func (outTxn *TxOutput) Print() {
	fmt.Printf("OUTPUT TXN ::\n Value : %d TO : %s\n", outTxn.Value, outTxn.ScriptPubKey)
}

