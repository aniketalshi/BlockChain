package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

const checksumBytes = 4
const version = 1

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallet() *Wallet {
	privKey, pubKey := NewKeyPair()
	return &Wallet{privKey, pubKey}
}

// NewKeyPair generates pair of private and associated public key
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	privateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)

	// generate public key by appending X and Y coordinates on curve
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return *privateKey, pubKey
}

// GenerateAddress is used of generating new addresses (Identity for transactions)
func (w Wallet) GenerateAddress() []byte {

	pubKeyHash := HashPubKey(w.PublicKey)
	payload := append([]byte{version}, pubKeyHash...)
	checkSum := GenerateChecksum(payload)
	address := append(payload, checkSum...)

	//address = Base58Encode(address)

	return address
}

//GenerateChecksum generates the checksum for given payload
func GenerateChecksum(block []byte) []byte {

	firstHash := sha256.Sum256(block)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:checksumBytes]
}

// HashPubKey generates hash of public key which will be used in constructing address
func HashPubKey(pubKey []byte) []byte {
	hashedKey := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()

	RIPEMD160Hasher.Write(hashedKey[:])
	publicHashKey := RIPEMD160Hasher.Sum(nil)

	return publicHashKey

}
