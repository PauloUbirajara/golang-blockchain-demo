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

func (bc *Blockchain) GetAllTransactions() []string {
	transactions := make([]string, 0)

	for _, b := range bc.Blocks {
		transactions = append(transactions, b.Content...)
	}

	return transactions
}

func (bc *Blockchain) NewBlock(difficulty int, content ...string) {
	DEFAULT_PREVIOUS_HASH := "secret"

	var lastBlock Block
	lastBlock.CurrentHash = DEFAULT_PREVIOUS_HASH

	if len(bc.Blocks) > 0 {
		lastBlock = bc.Blocks[len(bc.Blocks)-1]
	}

	validContent := make([]string, 0)
	userWallets := getUserWalletsFromTransactions(bc.GetAllTransactions())

	for _, trx := range content {
		if isValidCreationTransaction(trx) {
			// Adicionar no validContent + Criar carteira com valor definido
			validContent = append(validContent, trx)
			runCreationTransaction(userWallets, trx)
		}

		if isValidTransferTransaction(userWallets, trx) {
			// Adicionar no validContent + Transferir valor
			validContent = append(validContent, trx)
			runTransferTransaction(userWallets, trx)
		}
	}

	if len(validContent) == 0 {
		return
	}

	merkelRoot := merkelTree(validContent)

	newBlock := Block{
		Index:        lastBlock.Index + 1,
		Content:      validContent,
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

func (bc *Blockchain) ValidateBlockSequence(currentBlock Block, nextBlock Block) bool {
	expectedHash := currentBlock.CurrentHash
	currentBlock.SearchHash(currentBlock.Difficulty)

	for i := range currentBlock.CurrentHash {
		if currentBlock.CurrentHash[i] != expectedHash[i] {
			return false
		}
	}

	return true
}

func (bc *Blockchain) Validate() bool {
	if len(bc.Blocks) == 0 {
		return false
	}

	for i := 0; i < len(bc.Blocks)-1; i++ {
		currentBlock := bc.Blocks[i]
		nextBlock := bc.Blocks[i+1]

		if currentBlock.CurrentHash != nextBlock.PreviousHash {
			return false
		}

		if !bc.ValidateBlockSequence(currentBlock, nextBlock) {
			return false
		}
	}

	return bc.ValidateUserWalletCreation()
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
