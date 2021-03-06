package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"bufio"
)

func parseFile(fileName string) []float64 {
	var eqParams []float64
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	stringData := strings.Replace(string(data), "\r", "\n", -1)
	fileLines := strings.Split(stringData, "\n")

	values := strings.Split(fileLines[0], " ")
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

	consoleReader := bufio.NewReader(os.Stdin)
	for len(eqParams) != 3 {
		fmt.Printf("%s = ", argNames[len(eqParams)])

		input, err := consoleReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		input = strings.Replace(input, "\n", "", -1)
		valFloat, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}

		if valFloat == 0 && len(eqParams) == 0 {
			fmt.Fprint(os.Stderr, "Error: a cannot be 0\n")
			continue
		}

		eqParams = append(eqParams, valFloat)
	}

	return eqParams
}

func solveEquation(eqParams []float64) []float64 {
	var eqRoots []float64
	eqDiscriminant := eqParams[1]*eqParams[1] - 4*eqParams[0]*eqParams[2]

	if eqDiscriminant == 0 {
		singleRoot := -eqParams[1] / (2 * eqParams[0])
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
