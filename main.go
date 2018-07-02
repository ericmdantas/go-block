package main

import (
	"fmt"

	"github.com/ericmdantas/go-block/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("wat 1")
	bc.AddBlock("wat 2")
	bc.AddBlock("wat 3")

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
