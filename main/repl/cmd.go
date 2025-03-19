package repl

import (
	"coopstools/brainf-k/main/runner"
	"coopstools/brainf-k/main/tokenize"
	"errors"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"os"
	"strings"
)

var Green = "\033[32m"
var White = "\033[97m"
var Red = "\033[31m"

type CLI struct {
	msg    string
	window int

	runner runner.Runner

	pointer int
}

func (cli *CLI) Run() {
	rl := NewReadline()
	for {
		rl.SetPrompt(cli.runner.Draw(cli.window) + "> ")
		line, err := rl.Readline()
		if errors.Is(err, readline.ErrInterrupt) {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		cli.handleCommand(line)
	}
}

func (cli *CLI) handleCommand(command string) {
	switch command {
	case "help":
		cli.printHelp()
	case "reset":
		cli.msg = fmt.Sprintf("%s10%s [0 0 %s0%s 0 0]", Red, White, Green, White)
		cli.runner.Reset()
	case "inc":
		if cli.window < 21 {
			cli.window += 1
		}
	case "dec":
		if cli.window > 3 {
			cli.window -= 1
		}
	case "exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		cli.scan(command)
	}
}

func (cli *CLI) printHelp() {
	fmt.Println("Welcome to the BrainF--k REPL")
	fmt.Println("  help - Shows this help message")
	fmt.Println("  reset - wipe the current stack and start over")
	fmt.Println("  inc - increases window size")
	fmt.Println("  dec - decreases window size")
	fmt.Println("  exit - Exits the CLI")
}

func (cli *CLI) scan(command string) {
	cmds := tokenize.Tokenize(command)
	cli.runner.Run(cmds)
}

func New() *CLI {
	return &CLI{
		msg:    "",
		window: 3,
		runner: runner.New(os.Stdout),
	}
}
