package compiler

import (
	"coopstools/brainf-k/main/tokenize"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed template.txt
var template string

var cmdLookup = map[tokenize.Op]string{
	tokenize.INC_IND: "movr();",
	tokenize.DEC_IND: "movl();",
	tokenize.INC_VAL: "inc();",
	tokenize.DEC_VAL: "dec();",
	tokenize.WR_OUT:  "out();",
}

func inject(cmds []tokenize.Cmd, depth int) string {
	var subfuncs []string
	code := ""
	subcount := 1
	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		if cmd.Op == tokenize.CTRL_JMP {
			subfunc := inject(cmds[i+1:i+cmd.Value], depth*10+subcount)
			subfuncs = append(subfuncs, subfunc)
			i = i + cmd.Value
			code = fmt.Sprintf("%s  while(ptr[i]!=0){f%d();}\n", code, depth*10+subcount)
			subcount += 1
			continue
		}
		if cmd.Op == tokenize.RD_IN {
			if cmd.Value != -1 {
				code = fmt.Sprintf("%s  set(%d);\n", code, cmd.Value)
				continue
			}
			code = fmt.Sprintf("%s  setFrom();\n", code)
			continue
		}
		if cmd.Op == tokenize.WR_DEBUG {
			code = fmt.Sprintf("%s  debug(%d);\n", code, cmd.Value)
		}
		if v, ok := cmdLookup[cmd.Op]; ok {
			code = fmt.Sprintf("%s  %s\n", code, v)
		}
	}
	subfuncsJoined := strings.Join(subfuncs, "")
	return fmt.Sprintf("%s\nvoid f%d() {\n%s}", subfuncsJoined, depth, code)
}

func BuildIntoC(cmds []tokenize.Cmd) string {
	code := inject(cmds, 0)
	return strings.Replace(template, "{{$funcs}}", code, 1)
}
