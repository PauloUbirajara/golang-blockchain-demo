package basic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func (bc *Blockchain) ValidateSingleTransaction(userWallets map[string]float64, transaction string) bool {
	transactionBetweenUsers := strings.Split(transaction, " > ")

	// Verificar se primeiro usuário previamente criado com carteira
	firstUser := transactionBetweenUsers[0]
	firstUserValue, firstUserExists := userWallets[firstUser]

	if !firstUserExists {
		return false
	}

	// Verificar se segundo usuário previamente criado com carteira
	secondUser := transactionBetweenUsers[2]
	_, secondUserExists := userWallets[secondUser]

	if !secondUserExists {
		return false
	}

	valueToTransfer, err := strconv.ParseFloat(transactionBetweenUsers[1], 64)

	// Verificar por erro ao converter valor em string para float64
	if err != nil {
		return false
	}

	// Verificar se transação transformará o primeiro usuário em saldo negativo
	if firstUserValue-valueToTransfer < 0 {
		return false
	}

	// Alterar, na referência da carteira dos usuários, o novo valor dos usuários da transação
	userWallets[firstUser] -= valueToTransfer
	userWallets[secondUser] += valueToTransfer

	return true
}

func (bc *Blockchain) ValidateUserTransactions(userWallets map[string]float64) bool {
	USER_TRANSACTION_REGEX := `^[a-zA-Z]+ > [0-9]+(\.[0-9]+)? > [a-zA-Z]+$`

	// Verificar se transação respeita padrão A > VALOR > B
	re, err := regexp.Compile(USER_TRANSACTION_REGEX)

	if err != nil {
		return false
	}

	for _, block := range bc.Blocks[1:] {
		for _, trx := range block.Content {
			matched := re.Match([]byte(trx))

			if !matched {
				return false
			}

			if !bc.ValidateSingleTransaction(userWallets, trx) {
				return false
			}
		}
	}

	return true
}

func (bc *Blockchain) ValidateUserWalletCreation() bool {
	// Definindo carteiras
	userWallets := make(map[string]float64)

	USER_WALLET_REGEX := `^[a-zA-Z]+ = [0-9]+$`

	re, err := regexp.Compile(USER_WALLET_REGEX)

	if err != nil {
		return false
	}

	for _, trx := range bc.Blocks[0].Content {
		// Verificar se criação de carteira segue padrão A = VALOR
		matched := re.Match([]byte(trx))

		if err != nil || !matched {
			return false
		}

		currentUserWallet := strings.Split(trx, " = ")

		user := currentUserWallet[0]

		// Verifica se carteira previamente criada
		_, userAlreadyCreated := userWallets[user]
		if userAlreadyCreated {
			return false
		}

		value, err := strconv.ParseFloat(currentUserWallet[1], 64)

		if err != nil {
			return false
		}

		// Cria nova carteira e atribui valor
		userWallets[user] = value
	}

	return bc.ValidateUserTransactions(userWallets)
}

func (bc *Blockchain) Validate() bool {
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
