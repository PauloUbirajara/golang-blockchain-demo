package basic

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

type Blockchain struct {
	Blocks []Block
}

func (bc *Blockchain) NewBlock(content string, difficulty int) {
	DEFAULT_PREVIOUS_HASH := "secret"

	var lastBlock Block
	lastBlock.CurrentHash = DEFAULT_PREVIOUS_HASH

	if len(bc.Blocks) > 0 {
		lastBlock = bc.Blocks[len(bc.Blocks)-1]
	}

	newBlock := Block{
		Index:        lastBlock.Index + 1,
		Content:      content,
		PreviousHash: lastBlock.CurrentHash,
	}

	newBlock.SearchHash(difficulty)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.Blocks {
		b.Print()
	}
}

func (bc *Blockchain) SaveToJSON(outputName string) {
	file, _ := json.MarshalIndent(bc.Blocks, "", "  ")
	_ = ioutil.WriteFile(outputName, file, fs.ModeAppend.Perm())
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

	var blocks []Block
	json.Unmarshal(byteValue, &blocks)

	bc.Blocks = blocks
	log.Default().Println("Blockchain carregada com sucesso")
}
