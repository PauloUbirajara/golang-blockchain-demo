package basic

import (
	"regexp"
	"strconv"
	"strings"
)

type TransactionRegex string

func isValidTransaction(regex TransactionRegex, transaction string) bool {
	re, err := regexp.Compile(string(regex))

	return err == nil && re.Match([]byte(transaction))
}

func getArgsFromCreationTransaction(transaction string) (string, float64, error) {
	userCreationArgs := strings.Split(transaction, " = ")
	userToCreate := userCreationArgs[0]
	startingValue, err := strconv.ParseFloat(userCreationArgs[1], 64)

	if err != nil {
		return "", 0, err
	}

	return userToCreate, startingValue, nil
}

func getArgsFromTransferTransaction(transaction string) (string, float64, string, error) {
	transactionBetweenUsers := strings.Split(transaction, " > ")

	valueToTransfer, err := strconv.ParseFloat(transactionBetweenUsers[1], 64)
	if err != nil {
		return "", 0, "", err
	}

	firstUser := transactionBetweenUsers[0]
	secondUser := transactionBetweenUsers[2]

	return firstUser, valueToTransfer, secondUser, nil
}

func isValidTransferTransaction(userWallets map[string]float64, transaction string) bool {
	if !isValidTransaction(USER_TRANSFER_REGEX, transaction) {
		return false
	}

	firstUser, valueToTransfer, secondUser, err := getArgsFromTransferTransaction(transaction)
	if err != nil {
		return false
	}

	firstUserValue, firstUserExists := userWallets[firstUser]
	if !firstUserExists {
		return false
	}

	if firstUserValue < valueToTransfer {
		return false
	}

	_, secondUserExists := userWallets[secondUser]
	return secondUserExists
}

func isValidCreationTransaction(transaction string) bool {
	if !isValidTransaction(USER_CREATION_REGEX, transaction) {
		return false
	}

	_, startingValue, err := getArgsFromCreationTransaction(transaction)
	if err != nil {
		return false
	}

	return startingValue >= 0
}

func getUserWalletsFromTransactions(transactions []string) map[string]float64 {
	userWallets := make(map[string]float64)

	for _, trx := range transactions {
		if isValidCreationTransaction(trx) {
			userToCreate, startingValue, err := getArgsFromCreationTransaction(trx)
			if err != nil {
				continue
			}

			_, userExists := userWallets[userToCreate]
			if userExists {
				continue
			}

			userWallets[userToCreate] = startingValue
		}
	}

	return userWallets
}

func runCreationTransaction(userWallets map[string]float64, transaction string) {
	userToCreate, startingValue, err := getArgsFromCreationTransaction(transaction)
	if err != nil {
		return
	}

	userWallets[userToCreate] = startingValue
}

func runTransferTransaction(userWallets map[string]float64, transaction string) {
	firstUser, valueToTransfer, secondUser, err := getArgsFromTransferTransaction(transaction)

	if err != nil {
		return
	}

	userWallets[firstUser] -= valueToTransfer
	userWallets[secondUser] += valueToTransfer
}
