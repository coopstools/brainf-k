package main

import (
	"coopstools/brainf-k/main/runner"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	clArgs := os.Args
	for i, arg := range clArgs {
		if strings.Contains(arg, "main.go") {
			clArgs = clArgs[i:]
		}
	}

	contents, err := readInCode(clArgs)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
	}

	clArgs, opts := extractOpts(clArgs)
	args := extractCodeArgs(clArgs)

	results, stack := runner.RunBF(contents, args...)
	if opts['d'] {
		fmt.Println(stack)
	}

	if opts['s'] {
		return
	}

	tf := func(vals []byte) string {
		return fmt.Sprintf("%v", vals)
	}
	if opts['c'] {
		tf = func(vals []byte) string {
			return string(vals)
		}
	}
	fmt.Println(tf(results))
}

func extractOpts(args []string) ([]string, map[rune]bool) {
	opts := make(map[rune]bool)
	for i := 0; i < len(args); i++ {
		if args[i][0] != '-' || len(args[i]) == 1 {
			continue
		}
		for _, ropt := range args[i] {
			opts[ropt] = true
		}
		if len(args)-1 == i {
			args = args[:i-1]
			continue
		}
		args = append(args[:i], args[i+1:]...)
	}
	return args, opts
}

func readInCode(clArgs []string) ([]byte, error) {
	if len(clArgs) <= 1 {
		return nil, errors.New("no filename supplied")
	}
	file, err := os.Open(clArgs[1])
	if err != nil {
		return nil, errors.New("could not open file: " + err.Error())
	}
	contents, err := io.ReadAll(file)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n", err.Error())
		return nil, errors.New("could not pull file contents: " + err.Error())
	}
	return contents, nil
}

func extractCodeArgs(clArgs []string) []byte {
	var args []byte
	if len(clArgs) > 2 {
		for _, arg := range os.Args[2:] {
			v, err := strconv.Atoi(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "\ncould not parse value into int (0 to 255): %s\n", err.Error())
			}
			if v < 0 || v > 255 {
				_, _ = fmt.Fprintf(os.Stderr, "\ncould not parse value into int (0 to 255): %d\n", v)
			}
			args = append(args, byte(v))
		}
	}
	return args
}
