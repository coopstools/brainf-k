package repl

import (
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

	stack stack

	pointer int
}

func (cli *CLI) Run() {
	rl := NewReadline()
	for {
		cli.redraw()
		rl.SetPrompt(cli.msg + "> ")
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
		cli.stack = stack{s: make([]byte, 10000), i: 20}
	case "inc":
		if cli.window < 21 {
			cli.window += 2
		}
	case "dec":
		if cli.window > 3 {
			cli.window -= 2
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
	compiled := Compile(command + " ") //dumb fix for making sure input buffer is flushed
	for i := 0; i < len(compiled); i++ {
		switch compiled[i].op {
		case INC_IND:
			cli.stack.incInd()
		case DEC_IND:
			cli.stack.decInd()
		case INC_VAL:
			cli.stack.incVal()
		case DEC_VAL:
			cli.stack.decVal()
		case RD_IN:
			cli.stack.in(byte(compiled[i].value))
		case RW_OUT:
			fmt.Printf("%d\n", cli.stack.out())
		case CTRL_JMP:
			if cli.stack.out() != 0 {
				continue
			}
			i = compiled[i].value
		case CTRL_RTN:
			if cli.stack.out() == 0 {
				continue
			}
			i = compiled[i].value
		}
	}
}

func (cli *CLI) redraw() {
	ind, halfWin := cli.stack.i, cli.window/2
	cli.msg = fmt.Sprintf("%s%3d%s [", Green, ind, White)
	if ind >= halfWin {
		for _, val := range cli.stack.s[ind-halfWin : ind] {
			cli.msg += fmt.Sprintf("%3d ", val)
		}
		cli.msg += fmt.Sprintf("%s%3d%s", Red, cli.stack.s[ind], White)
		for _, val := range cli.stack.s[ind+1 : ind+halfWin+1] {
			cli.msg += fmt.Sprintf(" %3d", val)
		}
		cli.msg += "]"
	}
}

func New() *CLI {
	return &CLI{
		msg:    "",
		window: 5,
		stack:  stack{s: make([]byte, 10000), i: 20},
	}
}
