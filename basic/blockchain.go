package basic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Blockchain struct {
	Blocks []Block
}

func (bc *Blockchain) NewBlock(difficulty int, content ...string) {
	DEFAULT_PREVIOUS_HASH := "secret"

	var lastBlock Block
	lastBlock.CurrentHash = DEFAULT_PREVIOUS_HASH

	if len(bc.Blocks) > 0 {
		lastBlock = bc.Blocks[len(bc.Blocks)-1]
	}

	merkelRoot := merkelTree(content)

	newBlock := Block{
		Index:        lastBlock.Index + 1,
		Content:      content,
		PreviousHash: lastBlock.CurrentHash,
		Timestamp:    time.Now(),
		MerkelRoot:   merkelRoot,
	}

	newBlock.SearchHash(difficulty)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.Blocks {
		b.Print()
	}
}

func (bc *Blockchain) Validate() bool {
	for i := 0; i < len(bc.Blocks)-1; i++ {
		currentBlock := bc.Blocks[i]

		expectedHash := currentBlock.CurrentHash
		currentBlock.SearchHash(currentBlock.Difficulty)

		for i := range currentBlock.CurrentHash {
			if currentBlock.CurrentHash[i] != expectedHash[i] {
				return false
			}
		}

		nextBlock := bc.Blocks[i+1]

		if currentBlock.CurrentHash != nextBlock.PreviousHash {
			return false
		}
	}

	return true
}

func (bc *Blockchain) SaveToJSON(outputName string) {
	file, _ := json.MarshalIndent(bc.Blocks, "", "  ")
	_ = ioutil.WriteFile(outputName, file, 0644)
}

func (bc *Blockchain) LoadFromJSON(inputName string) {
	jsonFile, err := os.Open(inputName)

	if err != nil {
		log.Panic("Não foi possível carregar blockchain a partir de JSON", err)
		return
	}

	log.Default().Println("Arquivo aberto sem erros")
	fmt.Println("Arquivo aberto sem erros")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &bc.Blocks)
	log.Default().Println("Blockchain carregada com sucesso")
}

func merkelTree(content []string) string {
	if len(content) == 1 {
		return HashFromTransactions(content[0])
	}

	if len(content)%2 != 0 {
		content = append(content, content[len(content)-1])
	}

	emptyArray := make([]string, 0)

	for i := 0; i < len(content); i += 2 {
		emptyArray = append(emptyArray, HashFromTransactions(content[i]+content[i+1]))
	}

	return merkelTree(emptyArray)
}
