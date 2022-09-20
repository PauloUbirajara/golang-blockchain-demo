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

func BadExample1(blockchain *b.Blockchain) {
	// Carteira não criada conforme padrão
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A != 10", "B = 20")
}

func BadExample2(blockchain *b.Blockchain) {
	// Carteira criada em duplicidade
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "A = 20")
}

func BadExample3(blockchain *b.Blockchain) {
	// Transação para uma carteira inexistente
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "B = 10")
	blockchain.NewBlock(DIFFICULTY, "A > 5 > C")
}

func BadExample4(blockchain *b.Blockchain) {
	// Transação causa valor negativo em carteira
	DIFFICULTY := 3
	blockchain.NewBlock(DIFFICULTY, "A = 10", "B = 10")
	blockchain.NewBlock(DIFFICULTY, "A > 20 > B")
}

func main() {
	var blockchain b.Blockchain

	// GoodExample1(&blockchain)
	// BadExample1(&blockchain)
	// BadExample2(&blockchain)
	// BadExample3(&blockchain)
	BadExample4(&blockchain)

	blockchain.PrintBlocks()
	fmt.Println("Validado", blockchain.Validate())
}
