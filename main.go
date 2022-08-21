package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Nounce       int
	Content      string
	Timestamp    time.Time
	PreviousHash string
	CurrentHash  string
}

func (b *Block) String() string {
	finalString := ""

	finalString += fmt.Sprintln("Conteúdo:", b.Content)
	finalString += fmt.Sprintln("Timestamp:", b.Timestamp)
	finalString += fmt.Sprintln("Nounce:", b.Nounce)

	return finalString
}

func HashToString(hash []byte) string {
	finalString := fmt.Sprintf("%x", hash)

	return finalString
}

func (b *Block) SearchHash(difficulty int) {
	b.Nounce = 1
	b.Timestamp = time.Now()

	for {
		currentHash := sha256.Sum256([]byte(b.String()))
		validHash := true
		fmt.Println(HashToString(currentHash[:]))

		for i, letter := range currentHash {
			if i >= difficulty {
				break
			}
			if letter != '0' {
				validHash = false
				break
			}
		}

		if validHash {
			b.CurrentHash = HashToString(currentHash[:])
			return
		}

		b.Nounce++
	}
}

func main() {
	b := Block{}
	fmt.Println(b.String())
	b.SearchHash(1)
	fmt.Println(b)
	// TODO Fazer um método para imprimir melhor o bloco
}
