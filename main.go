package main

import (
	"fmt"
	b "golang-blockchain-demo/basic"
)

func main() {
	var blockchain b.Blockchain
	// blockchain.LoadFromJSON("blocks.json")

	// Adicionar N blocos, incluindo o genesis se "blockchain" iniciar vazio
	DIFFICULTY := 3
	BLOCK_COUNT := 3
	for i := 0; i < BLOCK_COUNT; i++ {
		blockchain.NewBlock([]string{"alo", "hello", "ola", "oi2"}, DIFFICULTY)
	}

	blockchain.PrintBlocks()
	fmt.Println("Validado", blockchain.Validate())
	// blockchain.SaveToJSON("blocks.json")
}
