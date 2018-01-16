package blockchain

import (
	"bytes"
	"math/big"
)

var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(address []byte) []byte {

	var result []byte

	base := big.NewInt(0).SetBytes(address)
	zero := big.NewInt(0)
	fiftyeight := big.NewInt(int64(len(base58Alphabet)))

	for base.Cmp(zero) != 0 {
		mod := &big.Int{}
		base.DivMod(base, fiftyeight, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	// if bitcoin address starts with 0x00
	for _, data := range address {
		if data == 0x00 {
			result = append([]byte{base58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for _, b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(base58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
