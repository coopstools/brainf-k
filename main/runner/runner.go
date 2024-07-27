package runner

import (
	"bytes"
	"fmt"
	"strconv"
)

func RunBF(code []byte, input ...byte) ([]byte, string) {
	buf := bytes.Buffer{}
	s := stack{s: make([]byte, 1000)}
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			s.incInd()
		case '<':
			if s.decInd() != nil {
				panic("index moved out of rang on op " + strconv.Itoa(i))
			}
		case '+':
			s.incVal()
		case '-':
			s.decVal()
		case '.':
			_ = buf.WriteByte(s.out())
		case ',':
			s.in(input[0])
			if len(input) > 1 {
				input = input[1:]
			}
		case '[':
			s.c += 1
			if s.val() != 0 {
				continue
			}
			si := i
			i = scanRight(i+1, code)
			if i < 0 {
				line := string(code[:si]) + "\x1B[41m" + string(code[si:si+1]) + "\x1B[0m" + "blue"
				fmt.Println(line)
				panic(fmt.Sprint("could not find matching bracket: ", line))
			}
		case ']':
			s.c += 1
			if s.val() == 0 {
				continue
			}
			i = scanLeft(i-1, code)
		case '#': //extra code for dumping stack state
			fmt.Println("dump", s.i, s.s[:17])
		}
	}
	return buf.Bytes(), fmt.Sprintf("%d %v %d", s.i, s.s[:17], s.c)
}

func scanLeft(index int, code []byte) int {
	initial := index
	count := 1
	for count > 0 && index >= 0 {
		switch code[index] {
		case '[':
			if count == 1 {
				return index
			}
			count--
		case ']':
			count++
		}
		index--
	}
	panic(fmt.Sprintf("could not find matching ] for symbol at %d", initial))
}

func scanRight(index int, code []byte) int {
	count := 1
	for count > 0 && index < len(code) {
		switch code[index] {
		case '[':
			count++
		case ']':
			if count == 1 {
				return index
			}
			count--
		}
		index++
	}
	return -1
}
