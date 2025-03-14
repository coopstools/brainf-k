package repl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompile(t *testing.T) {
	cmd := "[[,28.][<+>-]]"
	compiled := Compile(cmd)
	assert.Equal(t, 12, len(compiled))
	for i, expected := range []Cmd{
		{op: CTRL_JMP, value: 11}, // [
		{op: CTRL_JMP, value: 4},  // [
		{op: RD_IN, value: 28},    // ,28
		{op: RW_OUT},              // .
		{op: CTRL_RTN, value: 1},  // ]
		{op: CTRL_JMP, value: 10}, // [
	} {
		assert.Equal(t, expected, compiled[i], "error on index %d", i)
	}
}
