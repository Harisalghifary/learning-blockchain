package main

// method add a new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	PreviousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// method to return the whole blockchain and add genesis
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
