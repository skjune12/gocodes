package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter String: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.ToLower(input)

	var greetings = []string{"hello", "hola", "bonjour"}

	for _, greeting := range greetings {
		if strings.Contains(input, greeting) {
			fmt.Println(strings.Contains(input, greeting))
			return
		}
	}
	fmt.Println(false)
}
