package basic

const (
	// A > VALOR > B
	USER_TRANSFER_REGEX TransactionRegex = `^[a-zA-Z]+ > [0-9]+(\.[0-9]+)? > [a-zA-Z]+$`

	// A = VALOR
	USER_CREATION_REGEX = `^[a-zA-Z]+ = [0-9]+$`
)