package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Dans ma valise")
	fmt.Println("--------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		fmt.Println("Dans ma valise, il y a...", text)
	}

}
