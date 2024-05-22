package main

import (
	"fmt"
)

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func camelCase(s string) int32 {
	result := int32(1)
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			result++
		}
	}
	return result

}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	result := camelCase(input)
	fmt.Printf("%d", result)
}
