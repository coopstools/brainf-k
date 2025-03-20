package tokenize

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Op int

const (
	INC_IND  Op = 0
	DEC_IND  Op = 1
	INC_VAL  Op = 2
	DEC_VAL  Op = 3
	RD_IN    Op = 4
	WR_OUT   Op = 5
	CTRL_JMP Op = 6
	CTRL_RTN Op = 7

	WR_DEBUG = 8
)

type Cmd struct {
	Op    Op
	Value int
}

type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

func Tokenize(rawCmd string) []Cmd {
	rawCmd += " "
	var cmds []Cmd
	pStack := stack{values: make([]int, 0)}
	var scan []rune
	var p int
	for _, r := range rawCmd {
		if scan != nil {
			if unicode.IsDigit(r) {
				scan = append(scan, r)
				continue
			}
			val := -1
			if len(scan) >= 1 {
				val, _ = strconv.Atoi(string(scan))
			}
			cmds = append(cmds, Cmd{Op: RD_IN, Value: val % 256})
			scan = nil
			p++
		}
		switch r {
		case '>':
			cmds = append(cmds, Cmd{Op: INC_IND})
			p++
		case '<':
			cmds = append(cmds, Cmd{Op: DEC_IND})
			p++
		case '+':
			cmds = append(cmds, Cmd{Op: INC_VAL})
			p++
		case '-':
			cmds = append(cmds, Cmd{Op: DEC_VAL})
			p++
		case ',':
			scan = []rune{}
		case '.':
			cmds = append(cmds, Cmd{Op: WR_OUT})
			p++
		case '[':
			cmds = append(cmds, Cmd{Op: CTRL_JMP})
			pStack.push(p)
			p++
		case ']':
			lastP := pStack.pop()
			cmds = append(cmds, Cmd{Op: CTRL_RTN, Value: p - lastP})
			cmds[lastP].Value = p - lastP
			p++
		case '#':
			if len(cmds) > 0 && cmds[len(cmds)-1].Op == WR_DEBUG {
				cmds[len(cmds)-1].Value += 1
				break
			}
			cmds = append(cmds, Cmd{Op: WR_DEBUG, Value: 1})
			p++
		}
	}
	// TODO: Add check for mismatch parens
	// TODO: Add proper error handling to avoid breaking REPL
	// TODO: Use compiler in core code
	// TODO: Add stack dumps to compiler
	// TODO: Add stack dumps to Readme
	return cmds
}

func CreateTokensFromFileName(fileName string) []Cmd {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("could not open file: ", fileName, err.Error())
		os.Exit(4)
	}
	return Tokenize(string(contents))
}
