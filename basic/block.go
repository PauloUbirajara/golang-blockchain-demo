package basic

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Content      []string
	CurrentHash  string
	Difficulty   int
	Index        int
	Nounce       int
	PreviousHash string
	Timestamp    time.Time
	MerkelRoot   string
}

func (b *Block) StringForSHA256() string {
	finalString := ""

	finalString += fmt.Sprintln("Index:", b.Index)
	finalString += fmt.Sprintln("Conteúdo:", b.Content)
	finalString += fmt.Sprintln("Timestamp:", b.Timestamp.Nanosecond())
	finalString += fmt.Sprintln("Nounce:", b.Nounce)
	finalString += fmt.Sprintln("Hash Anterior:", b.PreviousHash)

	return finalString
}

func (b *Block) Print() {
	finalString := ""

	finalString += fmt.Sprintln("Index:", b.Index)
	finalString += fmt.Sprintln("Nounce:", b.Nounce)
	finalString += fmt.Sprintln("Conteúdo:", "[", strings.Join(b.Content, ", "), "]")
	finalString += fmt.Sprintln("Timestamp:", b.Timestamp)
	finalString += fmt.Sprintln("Hash atual:", b.CurrentHash)
	finalString += fmt.Sprintln("Hash anterior:", b.PreviousHash)
	finalString += fmt.Sprintln("Dificuldade:", b.Difficulty)
	finalString += fmt.Sprintln("MerkelRoot:", b.MerkelRoot)

	fmt.Println(finalString)
}

func (b *Block) HashFromContent() string {
	currentHash := sha256.Sum256([]byte(b.StringForSHA256()))
	hashString := HashToString(currentHash[:])

	return hashString
}

func (b *Block) SearchHash(difficulty int) {
	b.Nounce = 1
	b.Difficulty = difficulty

	for {
		hashString := b.HashFromContent()

		if CheckIfValidHash(hashString, difficulty) {
			b.CurrentHash = hashString
			return
		}

		b.Nounce++
	}
}
