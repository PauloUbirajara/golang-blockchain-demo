package main

import (
	"fmt"
	b "golang-blockchain-demo/basic"
)

func GoodExample1(blockchain *b.Blockchain) {
	// Carteiras criadas corretamente e transações respeitam limites de conta
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "B = 20")
	blockchain.NewBlock(DIFFICULTY, "A > 10 > B")
	blockchain.NewBlock(DIFFICULTY, "B > 2 > A", "A > 1 > B", "B > 9 > A")
}

func GoodExample2(blockchain *b.Blockchain) {
	// Carteiras criadas antes e depois do gênesis
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10")
	blockchain.NewBlock(DIFFICULTY, "B = 20")
}

func BadExample1(blockchain *b.Blockchain) {
	// Carteira não criada conforme padrão
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A != 10", "B = 20")
}

func BadExample2(blockchain *b.Blockchain) {
	// Transação para uma carteira inexistente
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "B = 10")
	blockchain.NewBlock(DIFFICULTY, "A > 5 > C")
}

func BadExample3(blockchain *b.Blockchain) {
	// Transação causa valor negativo em carteira
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "B = 10")
	blockchain.NewBlock(DIFFICULTY, "A > 20 > B")
}

func main() {
	var blockchain b.Blockchain

	// GoodExample1(&blockchain)
	GoodExample2(&blockchain)
	// BadExample1(&blockchain)
	// BadExample2(&blockchain)
	// BadExample3(&blockchain)

	blockchain.PrintBlocks()
	fmt.Println("Validado", blockchain.Validate())
}
