package repl

import "github.com/chzyer/readline"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("help"),
	readline.PcItem("refresh"),
	readline.PcItem("inc"),
	readline.PcItem("dec"),
	readline.PcItem("exit"),
)

func NewReadline() *readline.Instance {
	rl, _ := readline.NewEx(&readline.Config{
		AutoComplete: completer,
	})
	return rl
}
