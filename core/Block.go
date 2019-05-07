package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  // index of the block
	Timestamp     int64  // timestamp  the block create time
	PrevBlockHash string // previous block hash
	Hash          string // current block hash

	Date string // Block Data  (trade data)
}

func calculatHash(b *Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Date
	HashInByte := sha256.Sum256([]byte(blockData))

	// array to slice -->  a[:]
	return hex.EncodeToString(HashInByte[:])
}

func GenerateNewBlock(b *Block, data string) *Block {
	newBlock := Block{}
	newBlock.Index = b.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = b.Hash
	newBlock.Date = data
	newBlock.Hash = calculatHash(&newBlock)

	return &newBlock
}

func GenerateGenesisBlock() *Block {
	dummyBlock := Block{}
	dummyBlock.Index = -1
	dummyBlock.Hash = ""
	return GenerateNewBlock(&dummyBlock, "Genesis Block")
}
