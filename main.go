package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type text")

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			p := Encode(text)
			fmt.Printf("%s is %s in pig latin.\n", text, p)
		}
	}
}
