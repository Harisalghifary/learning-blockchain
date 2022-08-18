// main package
package main

// import package
import (
	"fmt"
)

func main() {
	fmt.Println("My first blockchain")

	// create a new blockchain
	CreateBlock("The header of the first block", "The body of the first block\n")

	// init blockchain
	bc := NewBlockchain()
	// create 2 block and add 2 transactions
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	// print blockchain
	for i, block := range bc.Blocks { // loop through blockchain
		fmt.Printf("Block ID: %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)             // print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.PrevBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.Data)
		fmt.Println()
	}
}
