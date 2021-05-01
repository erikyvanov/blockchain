package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Segundo bloque")
	bc.AddBlock("Tercer bloque")
	bc.AddBlock("Cuarto bloque")

	for _, block := range bc.blocks {
		fmt.Printf("Hash: %x\n\n", block.Hash)
		fmt.Printf("Hash anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
