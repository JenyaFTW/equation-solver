package main

import (
	"os"
	"fmt"
)

func main() {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	fmt.Println(fileName)
}
