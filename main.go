package main

import (
	"fmt"
	b "golang-blockchain-demo/basic"
)

func main() {
	DIFFICULTY := 3
	BLOCK_COUNT := 2

	var blockchain b.Blockchain
	blockchain.LoadFromJSON("blocks.json")

	// Adicionar N blocos, incluindo o genesis se "blockchain" iniciar vazio
	for i := 0; i < BLOCK_COUNT; i++ {
		blockContent := fmt.Sprintf("Bloco %d", i+1)
		blockchain.NewBlock(blockContent, DIFFICULTY)
	}

	blockchain.PrintBlocks()
	fmt.Println("Validado", blockchain.Validate())
	// blockchain.SaveToJSON("blocks.json")
}
