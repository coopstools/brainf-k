package repl

import (
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
	RW_OUT   Op = 5
	CTRL_JMP Op = 6
	CTRL_RTN Op = 7
)

type Cmd struct {
	op    Op
	value int
}

type Cmds []Cmd

func Compile(rawCmd string) Cmds {
	var cmds Cmds
	var pStack []int
	var scan []rune
	var p int
	for _, r := range rawCmd {
		if scan != nil {
			if unicode.IsDigit(r) {
				scan = append(scan, r)
				continue
			}
			val, _ := strconv.Atoi(string(scan))
			cmds = append(cmds, Cmd{op: RD_IN, value: val % 256})
			scan = nil
			p++
		}
		switch r {
		case '>':
			cmds = append(cmds, Cmd{op: INC_IND})
			p++
		case '<':
			cmds = append(cmds, Cmd{op: DEC_IND})
			p++
		case '+':
			cmds = append(cmds, Cmd{op: INC_VAL})
			p++
		case '-':
			cmds = append(cmds, Cmd{op: DEC_VAL})
			p++
		case ',':
			scan = []rune{}
		case '.':
			cmds = append(cmds, Cmd{op: RW_OUT})
			p++
		case '[':
			cmds = append(cmds, Cmd{op: CTRL_JMP})
			pStack = append(pStack, p)
			p++
		case ']':
			last_p := pStack[len(pStack)-1]
			pStack = pStack[:len(pStack)-1]
			cmds = append(cmds, Cmd{op: CTRL_RTN, value: last_p})
			cmds[last_p].value = p
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
