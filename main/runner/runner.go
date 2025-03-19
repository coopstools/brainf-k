package runner

import (
	"bytes"
	"coopstools/brainf-k/main/tokenize"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var Green = "\033[32m"
var White = "\033[97m"
var Red = "\033[31m"

type Runner struct {
	s   *stack
	out io.Writer
}

func (r Runner) Run(cmds []tokenize.Cmd) {
	for i := 0; i < len(cmds); i++ {
		switch cmds[i].Op {
		case tokenize.INC_IND:
			r.s.incInd()
		case tokenize.DEC_IND:
			if r.s.decInd() != nil {
				panic("index moved out of rang on op " + strconv.Itoa(i))
			}
		case tokenize.INC_VAL:
			r.s.incVal()
		case tokenize.DEC_VAL:
			r.s.decVal()
		case tokenize.RD_IN:
			if cmds[i].Value == -1 {
				panic("non immediate reference not allowed in this context; please supply value inline, ie `,123`")
			}
			r.s.in(byte(cmds[i].Value))
		case tokenize.RW_OUT:
			_, _ = fmt.Fprintf(r.out, "%d", r.s.out())
		case tokenize.CTRL_JMP:
			if r.s.out() == 0 {
				i += cmds[i].Value
			}
		case tokenize.CTRL_RTN:
			if r.s.out() != 0 {
				i -= cmds[i].Value
			}
		case tokenize.RW_DEBUG:
			_, _ = fmt.Fprintf(r.out, "%s\n", r.Draw(cmds[i].Value))
		}
	}
}

func (r Runner) Reset() {
	r.s.s = make([]byte, 16328)
}

func (r Runner) Draw(width int) string {
	center := r.s.i
	if center < width {
		center = width
	}
	svs := make([]string, width*2+1)
	for j, v := range r.s.s[center-width : center+width+1] {
		svs[j] = fmt.Sprintf("%2d", v)
	}
	indexLoc := r.s.i
	if r.s.i > width {
		indexLoc = width
	}
	svs[indexLoc] = fmt.Sprintf("%s%s%s", Red, svs[indexLoc], White)
	return fmt.Sprintf("%s%d%s [%s]",
		Red, r.s.i, White, strings.Join(svs, " "))
}

func New(out io.Writer) Runner {
	return Runner{
		s:   &stack{s: make([]byte, 16328)},
		out: out,
	}
}

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
			dump(s)
			// fmt.Println("dump", s.i, s.s[:17])
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

func dump(s stack) {
	dumpSize := 17
	svs := make([]string, dumpSize)
	for i, v := range s.s[:dumpSize] {
		svs[i] = fmt.Sprintf("%2d", v)
	}
	stackOut := "[" + strings.Join(svs, " ") + "]"
	fmt.Println("dump", s.i, stackOut)
}
