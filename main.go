package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Too few arguments!")
		os.Exit(1)
	}
	hexString := args[1]
	fmt.Println("CONVERTING HEX: ", hexString)

	cleanedHexString := cleanHex(hexString)
	fmt.Println(cleanedHexString)
	// Decode the hex string to bytes
	bytes, err := hex.DecodeString(cleanedHexString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Convert bytes to string
	fmt.Println("STRING:", string(bytes))

}

func cleanHex(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	// Remove "0x" or "0X" prefix if present
	if len(s) > 1 && (s[:2] == "0x" || s[:2] == "0X") {
		s = s[2:]
	}
	return s
}
