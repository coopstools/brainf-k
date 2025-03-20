package runner

import (
	"coopstools/brainf-k/main/tokenize"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Green = "\033[32m"
var White = "\033[97m"
var Red = "\033[31m"

type Runner struct {
	s *stack

	inputs []byte

	stdout *os.File
	stderr *os.File
}

func (r *Runner) Run(cmds []tokenize.Cmd) {
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
			if cmds[i].Value != -1 {
				r.s.in(byte(cmds[i].Value))
				continue
			}
			if len(r.inputs) == 0 {
				panic("non immediate reference not allowed in this context; please supply value inline, ie `,123`")
			}
			r.s.in(r.inputs[0])
			r.inputs = r.inputs[1:]
		case tokenize.WR_OUT:
			out := r.stdout
			if r.s.s[0] != 0 {
				out = r.stderr
			}
			switch r.s.s[1] {
			case 0:
				_, _ = fmt.Fprintf(out, "%d", r.s.out())
			case 1:
				_, _ = fmt.Fprintf(out, "%x", r.s.out())
			default:
				_, _ = fmt.Fprintf(out, "%c", r.s.out())
			}
		case tokenize.CTRL_JMP:
			if r.s.out() == 0 {
				i += cmds[i].Value
			}
		case tokenize.CTRL_RTN:
			if r.s.out() != 0 {
				i -= cmds[i].Value
			}
		case tokenize.WR_DEBUG:
			out := r.stdout
			if r.s.s[2] != 0 {
				out = r.stderr
			}
			_, _ = fmt.Fprintf(out, "%s\n", r.Draw(cmds[i].Value, r.s.s[3]))
		}
	}
	fmt.Fprintf(r.stdout, "\n")
}

func (r *Runner) Reset() {
	r.s.s = make([]byte, 16328)
}

func (r *Runner) Draw(width int, repSelect byte) string {
	rep := "%2d"
	if repSelect == 1 {
		rep = "%c"
	} else if repSelect == 2 {
		rep = "0x%x"
	}
	center := r.s.i
	if center < width {
		center = width
	}
	svs := make([]string, width*2+1)
	for j, v := range r.s.s[center-width : center+width+1] {
		svs[j] = fmt.Sprintf(rep, v)
	}
	indexLoc := r.s.i
	if r.s.i > width {
		indexLoc = width
	}
	svs[indexLoc] = fmt.Sprintf("%s%s%s", Red, svs[indexLoc], White)
	return fmt.Sprintf("%s%d%s [%s]",
		Red, r.s.i, White, strings.Join(svs, " "))
}

func (r *Runner) SetInputs(inputs []byte) *Runner {
	r.inputs = inputs
	return r
}

func New(stdout *os.File, stderr *os.File) *Runner {
	return &Runner{
		s:      &stack{s: make([]byte, 16328), i: 4},
		stdout: stdout,
		stderr: stderr,
	}
}
