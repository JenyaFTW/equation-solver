package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(fileName string) []float64 {
	var eqParams []float64
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	values := strings.Split(string(data), " ")
	if len(values) != 3 {
		fmt.Fprint(os.Stderr, "Error: Invalid file format, expected 3 values\n")
		os.Exit(1)
	}

	for _, val := range values {
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		if valFloat == 0 {
			fmt.Fprint(os.Stderr, "Value cannot be 0\n")
			os.Exit(1)
		}

		eqParams = append(eqParams, valFloat)
	}
	return eqParams
}

func main() {
	var eqParams []float64
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	if fileName != "" {
		eqParams = parseFile(fileName)
	} else {

	}

	fmt.Printf("Equation is: (%f) x^2 + (%f) x + (%f) = 0", eqParams[0], eqParams[1], eqParams[2])
}
