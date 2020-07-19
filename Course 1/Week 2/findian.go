package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // Prompts for user input
	fmt.Print("Enter text: ")           // Asks user to provide input
	text, _ := reader.ReadString('\n')  // Grabs input in text variable
	fmt.Println(text)                   // Prints text variable
	resStr := strings.ToLower(text)     // Returns string all lower

	if strings.Contains(resStr, "i") && strings.Contains(resStr, "a") && strings.Contains(resStr, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
