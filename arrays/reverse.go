package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ""
	fmt.Printf("Enter a string : ")

	// Scanner is used to parse the entire line, rather than to terminate reading at space.
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input = scanner.Text()
	fmt.Printf("Reversed string is: '%s'\n", reverse(input))
}

func reverse(input string) string {
	length := len(input)
	if length <= 2 {
		return input
	}

	t := ""
	for i := length - 1; i >= 0; i-- {
		t += string(input[i])
	}

	return t
}