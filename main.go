package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Content      string
	CurrentHash  string
	Index        int
	Nounce       int
	PreviousHash string
	Timestamp    time.Time
}

func (b *Block) StringForSHA256() string {
	finalString := ""

	finalString += fmt.Sprintln("Conteúdo:", b.Content)
	finalString += fmt.Sprintln("Timestamp:", b.Timestamp)
	finalString += fmt.Sprintln("Nounce:", b.Nounce)

	return finalString
}

func (b *Block) Print() {
	finalString := ""

	finalString += fmt.Sprintln("Index:", b.Index)
	finalString += fmt.Sprintln("Nounce:", b.Nounce)
	finalString += fmt.Sprintln("Conteúdo:", b.Content)
	finalString += fmt.Sprintln("Timestamp:", b.Timestamp)
	finalString += fmt.Sprintln("Hash atual:", b.CurrentHash)
	finalString += fmt.Sprintln("Hash anterior:", b.PreviousHash)

	fmt.Println(finalString)
}

func HashToString(hash []byte) string {
	finalString := fmt.Sprintf("%x", hash)

	return finalString
}

func CheckIfValidHash(hashString string, expectedZeros int) bool {
	zeroCount := 0

	for _, letter := range hashString {
		if letter != '0' {
			break
		}

		zeroCount++

		if zeroCount > expectedZeros {
			return false
		}
	}

	return zeroCount == expectedZeros
}

func (b *Block) SearchHash(difficulty int) {
	b.Nounce = 1
	b.Timestamp = time.Now()

	for {
		currentHash := sha256.Sum256([]byte(b.StringForSHA256()))
		hashString := HashToString(currentHash[:])

		println(hashString)

		if CheckIfValidHash(hashString, difficulty) {
			b.CurrentHash = hashString
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
