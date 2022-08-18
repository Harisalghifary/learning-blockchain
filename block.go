package main

import (
	"bytes"         // convert data into byte
	"crypto/sha256" // hash data
	"fmt"
	"strconv" // convert data into string
	"time"
)

func CreateBlock(header string, body string) {
	fmt.Println(header)
	fmt.Println(body)
}

// method attach to model Block
// set hash to block
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PrevBlockHash, block.Data}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]

}

// create new block generate and return new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{[]byte{}, prevBlockHash, time.Now().Unix(), []byte(data)}
	block.SetHash() // set hash to block
	return block
}

// create new genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
