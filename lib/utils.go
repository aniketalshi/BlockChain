package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"strconv"
)

//IntToHex converts given integer to hex
func IntToHex(num int64) []byte {
	return []byte(strconv.FormatInt(num, 16))
}

// DeserializeBlock deserializes given bytes into block struct
func DeserializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	if err := decoder.Decode(&block); err != nil {
		log.Fatal("DeserializeBlock :", err)
	}
	return &block
}
