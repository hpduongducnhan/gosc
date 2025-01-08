package gosc

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

func HashString(s string) string {
	// Create a new SHA-256 hash
	hasher := sha256.New()

	// Write the input string as bytes
	hasher.Write([]byte(s))

	// Compute the hash and convert it to a hexadecimal string
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func WriteToFile(exception string, fileName string, folder string) {
	if folder == "" {
		folder = "exceptions"
	}
	// create new folder if not exists
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0755)
	}

	// write exceptions to files
	file, err := os.OpenFile("exceptions/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %s", err)
	}
	defer file.Close()

	if _, err := file.WriteString(exception + "\n"); err != nil {
		log.Printf("Error writing to file: %s", err)
	}
}
