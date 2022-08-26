package basic

import "fmt"

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
