package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockChain() *Blockchain{
	genesisBlock := GenerateGenesisBlock()
	bc := &Blockchain{}
	bc.AppendBlock(genesisBlock)
	return bc
}

func (bc *Blockchain)SendDate(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(lastBlock, data)
	bc.AppendBlock(newBlock)
}

func (bc *Blockchain) AppendBlock (newBlock *Block) {
	if len(bc.Blocks) == 0 || isValid(newBlock, bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatalln("invalid block")
	}
	
}

func isValid(newBlock, oldBlock *Block) bool {
	if newBlock.Index != oldBlock.Index + 1 {
		return false
	}
	
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	
	if newBlock.Hash != calculatHash(newBlock) {
		return false
	}
	
	return true
}

func (bc *Blockchain) Print() {
	for _, b := range bc.Blocks {
		fmt.Printf("Index: %d\n", b.Index)
		fmt.Printf("Prev.Hash: %s\n", b.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", b.Hash)
		fmt.Printf("Curr.Data: %s\n", b.Date)
		fmt.Printf("Timestamp: %d\n", b.Timestamp)
		fmt.Println()
	}
}