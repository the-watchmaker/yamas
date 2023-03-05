package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args

	if len(args) == 3 {
		
		cmd := args[1]
		fileName :=	args[2]

		switch cmd {
			case "run":
				RunYamas(fileName)
				break
		}
		
	} else {
		fmt.Println("Please provide two arguments")
	}
}