package main

// block data structure
type Block struct {
	MyBlockHash   []byte
	PrevBlockHash []byte
	Timestamp     int64
	Data          []byte
}

// blockchain data structure
type Blockchain struct {
	Blocks []*Block
}
