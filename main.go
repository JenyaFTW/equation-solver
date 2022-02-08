package main

import (
	"fmt"
	"math"
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

	for idx, val := range values {
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		if valFloat == 0 && idx == 0 {
			fmt.Fprint(os.Stderr, "Error: a cannot be 0\n")
			os.Exit(1)
		}

		eqParams = append(eqParams, valFloat)
	}
	return eqParams
}

func parseStdIn() []float64 {
	var eqParams []float64

	argNames := []string{"a", "b", "c"}
	for idx, argName := range argNames {
		var input float64
		fmt.Printf("%s = ", argName)
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		if input == 0 && idx == 0 {
			fmt.Fprint(os.Stderr, "Error: Value cannot be 0\n")
			os.Exit(1)
		}

		eqParams = append(eqParams, input)
	}

	return eqParams
}

func solveEquation(eqParams []float64) []float64 {
	var eqRoots []float64
	eqDiscriminant := eqParams[1]*eqParams[1] - 4*eqParams[0]*eqParams[2]

	if eqDiscriminant == 0 {
		singleRoot := -eqParams[1] / 2 * eqParams[0]
		eqRoots = append(eqRoots, singleRoot)
	} else if eqDiscriminant > 0 {
		firstRoot := (-eqParams[1] + math.Sqrt(eqDiscriminant)) / (2 * eqParams[0])
		secondRoot := (-eqParams[1] - math.Sqrt(eqDiscriminant)) / (2 * eqParams[0])
		eqRoots = append(eqRoots, firstRoot, secondRoot)
	}

	return eqRoots
}

func main() {
	var eqParams []float64
	var eqRoots []float64

	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	if fileName != "" {
		eqParams = parseFile(fileName)
	} else {
		eqParams = parseStdIn()
	}

	eqRoots = solveEquation(eqParams)

	fmt.Printf("Equation is: (%f) x^2 + (%f) x + (%f) = 0\n", eqParams[0], eqParams[1], eqParams[2])
	fmt.Printf("There are %d root(s)\n", len(eqRoots))

	for idx, val := range eqRoots {
		fmt.Printf("x%d = %f\n", idx+1, val)
	}
}
