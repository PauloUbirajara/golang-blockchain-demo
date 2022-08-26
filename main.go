package main

import (
	"encoding/json"
	"fmt"
	b "golang-blockchain-demo/basic"
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

	blocks = readJsonFile();
	if blocks == nil {
		blocks = make([]b.Block, 0)

		// Primeiro bloco - Gênesis
		blocks = append(blocks, b.Genesis("Primeiro bloco", DIFFICULTY))
	}

	// Resto dos blocos
	for i := 0; i < BLOCK_COUNT; i++ {
		blockContent := fmt.Sprintf("Bloco %d", i)
		blocks = append(blocks, b.NewBlock(blocks, blockContent, DIFFICULTY))
	}

	for _, b := range blocks {
		b.Print()
	}

	file, _ := json.MarshalIndent(blocks, "", "  ")
	_ = ioutil.WriteFile("blocks.json", file, 0644);
}
