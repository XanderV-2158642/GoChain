package main

import (
	"fmt"
)

func main() {
	newblockchain := NewBlockChain()

	newblockchain.AddBlock("First Block")
	newblockchain.AddBlock("Second Block")

	for i, block := range newblockchain.Blocks {
		fmt.Printf("Block ID: %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.AllData)
		fmt.Printf("Hash: %x\n", block.MyBlockHash)
		fmt.Println()
	}
}
