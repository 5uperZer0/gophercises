package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func caesarCipher(s string, k int32) string {
	var outputBuilder strings.Builder
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			outputBuilder.WriteRune(((int32(ch) - int32('a') + k) % 26) + int32('a'))
		} else if ch >= 'A' && ch <= 'Z' {
			outputBuilder.WriteRune(((int32(ch) - int32('A') + k) % 26) + int32('A'))
		} else {
			outputBuilder.WriteRune(ch)
		}
	}
	return outputBuilder.String()
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plaintext: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Print("Enter an integer offset: ")
	k_str, _ := reader.ReadString('\n')
	k_str = strings.TrimSpace(k_str)
	k, _ := strconv.ParseInt(k_str, 10, 32)

	c := caesarCipher(s, int32(k))

	fmt.Printf("Original: %s\nCiphertext: %s\n", s, c)
}
