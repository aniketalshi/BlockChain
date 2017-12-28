package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

const targetbits = 24

// ProofOfWork struct which helps in boilerplate code for proof of work generation and validation
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork is a constructor
func NewProofOfWork(b *Block) *ProofOfWork {

	tt := big.NewInt(1)
	tt.Lsh(tt, uint(256-targetbits))
	pow := &ProofOfWork{b, tt}
	return pow
}

// PrepareData combines all data blocks and merges them into byte
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.Data,
			pow.block.PrevBlockHash,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetbits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run runs and caclulates hash. returns the nonce and corresponding hash value
func (pow *ProofOfWork) Run() (int, []byte) {

	var hashInt big.Int
	var hash [32]byte
	maxNonce := math.MaxInt64
	nonce := 0

	for nonce < maxNonce {

		mergedBytes := pow.PrepareData(nonce)
		hash = sha256.Sum256(mergedBytes)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	return nonce, hash[:]
}

// Validate will calculate and copare hash.
func (pow *ProofOfWork) Validate() bool {
	mergedBytes := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(mergedBytes)

	var hashInt big.Int
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1 {
		return true
	}
	return false
}
