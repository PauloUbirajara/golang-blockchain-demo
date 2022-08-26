package main

import (
	"encoding/json"
	"fmt"
	b "golang-blockchain-demo/basic"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func ReadJSONFile() []b.Block {
	jsonFile, err := os.Open("blocks.json")

	if err != nil {
		log.Println(err)
		return make([]b.Block, 0)
	}

	fmt.Println("File Opened succesfully!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var blocks []b.Block
	json.Unmarshal(byteValue, &blocks)

	return blocks
}

func main() {
	DIFFICULTY := 3
	BLOCK_COUNT := 25
	var blocks []b.Block

	blockchain := b.Blockchain{
		Blocks: ReadJSONFile(),
	}

	// Adicionar N blocos, incluindo o genesis se "blockchain" iniciar vazio
	for i := 0; i < BLOCK_COUNT; i++ {
		blockContent := fmt.Sprintf("Bloco %d", i+1)
		blockchain.NewBlock(blockContent, DIFFICULTY)
	}

	blockchain.PrintBlocks()

	file, _ := json.MarshalIndent(blockchain.Blocks, "", "  ")
	_ = ioutil.WriteFile("blocks.json", file, fs.ModeAppend.Perm())
}
