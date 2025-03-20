package compiler

import (
	"coopstools/brainf-k/main/tokenize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestInject(t *testing.T) {
	cmds := []tokenize.Cmd{
		{Op: tokenize.INC_IND}, {Op: tokenize.DEC_IND}, {Op: tokenize.INC_VAL}, {Op: tokenize.DEC_VAL},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body1, output[start+21:stop])
}

var body1 = `
void f0() {
  movr();
  movl();
  inc();
  dec();
}
`

func TestInject_flowControl_multiJmp(t *testing.T) {
	// ><[+<]>[+<]-
	cmds := []tokenize.Cmd{
		{Op: tokenize.INC_IND}, {Op: tokenize.DEC_IND},
		{Op: tokenize.CTRL_JMP, Value: 3}, // JMP 1
		{Op: tokenize.INC_VAL}, {Op: tokenize.DEC_IND},
		{Op: tokenize.CTRL_RTN, Value: 3},
		{Op: tokenize.INC_IND},
		{Op: tokenize.CTRL_JMP, Value: 3}, // JMP 2
		{Op: tokenize.INC_VAL}, {Op: tokenize.DEC_IND},
		{Op: tokenize.CTRL_RTN, Value: 3},
		{Op: tokenize.DEC_VAL},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body2, output[start+21:stop])
}

// ><[+<]>[+<]-
var body2 = `
void f1() {
  inc();
  movl();
}
void f2() {
  inc();
  movl();
}
void f0() {
  movr();
  movl();
  while(ptr[i]!=0){f1();}
  movr();
  while(ptr[i]!=0){f2();}
  dec();
}
`

func TestInject_flowControl_subJump(t *testing.T) {
	// >[>[-<]<]>[-]+
	cmds := []tokenize.Cmd{
		{Op: tokenize.INC_IND},
		{Op: tokenize.CTRL_JMP, Value: 7},
		{Op: tokenize.INC_IND},
		{Op: tokenize.CTRL_JMP, Value: 3},
		{Op: tokenize.DEC_VAL},
		{Op: tokenize.DEC_IND},
		{Op: tokenize.CTRL_RTN, Value: 3},
		{Op: tokenize.DEC_IND},
		{Op: tokenize.CTRL_RTN, Value: 7},
		{Op: tokenize.INC_IND},
		{Op: tokenize.CTRL_JMP, Value: 2},
		{Op: tokenize.DEC_VAL},
		{Op: tokenize.CTRL_RTN, Value: 2},
		{Op: tokenize.INC_VAL},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body3, output[start+21:stop])
}

// >[>[-<]<]>[-]+
var body3 = `
void f11() {
  dec();
  movl();
}
void f1() {
  movr();
  while(ptr[i]!=0){f11();}
  movl();
}
void f2() {
  dec();
}
void f0() {
  movr();
  while(ptr[i]!=0){f1();}
  movr();
  while(ptr[i]!=0){f2();}
  inc();
}
`

func TestInject_readIn(t *testing.T) {
	cmds := []tokenize.Cmd{
		{Op: tokenize.INC_IND}, {Op: tokenize.RD_IN, Value: 28}, {Op: tokenize.INC_IND}, {Op: tokenize.RD_IN, Value: -1}, {Op: tokenize.INC_IND},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body4, output[start+21:stop])
}

var body4 = `
void f0() {
  movr();
  set(28);
  movr();
  setFrom();
  movr();
}
`

func TestInject_moreFlorCtrl(t *testing.T) {
	cmds := []tokenize.Cmd{
		{Op: tokenize.CTRL_JMP, Value: 6},
		{Op: tokenize.INC_IND},
		{Op: tokenize.RD_IN, Value: 28},
		{Op: tokenize.INC_IND},
		{Op: tokenize.RD_IN, Value: -1},
		{Op: tokenize.INC_IND},
		{Op: tokenize.CTRL_RTN, Value: 6},
	}

	output := InjectTokensAsCode(cmds)
	start := strings.Index(output, "// start custom code")
	stop := strings.Index(output, "// stop custom code")
	require.NotEqual(t, -1, start, "missing beginning string")
	require.NotEqual(t, -1, stop, "missing ending string")
	assert.Equal(t, body4, output[start+21:stop])
}

var body5 = `
void f0() {
  movr();
  set(28);
  movr();
  setFrom();
  movr();
}
`
