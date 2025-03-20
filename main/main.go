package main

import (
	compiler "coopstools/brainf-k/main/compile"
	"coopstools/brainf-k/main/repl"
	"coopstools/brainf-k/main/runner"
	"coopstools/brainf-k/main/tokenize"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		repl.New().Run()
		return
	}

	clArgs := os.Args[1:]

	// Check for compile flag
	compileMode := false
	if clArgs[0] == "-c" {
		compileMode = true
		clArgs = clArgs[1:]
	}

	// Need filename argument
	if len(clArgs) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No filename provided")
		return
	}

	// Read input file
	contents, err := readInCode(clArgs[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if compileMode {
		// Compile mode - output C file
		tokens := tokenize.Tokenize(string(contents))
		cCode := compiler.BuildIntoC(tokens)
		outFile := filepath.Base(clArgs[0])
		outFile = outFile[:len(outFile)-len(filepath.Ext(outFile))] + ".c"
		err = os.WriteFile(outFile, []byte(cCode), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error writing output file:", err)
			return
		}
		return
	}
	// take the remaining arguments and parse out all numbers; delimiter can be spaces or commas or both
	var inputs []byte
	for _, arg := range clArgs[1:] {
		numStrings := strings.Split(arg, ",")
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: Invalid number:", numString)
				return
			}
			inputs = append(inputs, byte(num))
		}
	}
	fmt.Println(inputs)
	// Interpret mode - run code directly
	results, _ := runner.RunBF(contents, inputs...)
	fmt.Print(string(results))
}

func readInCode(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("could not open file: " + err.Error())
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("could not read file contents: " + err.Error())
	}
	return contents, nil
}
