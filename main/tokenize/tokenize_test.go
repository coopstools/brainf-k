package tokenize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	testCode := ">,28[<+>->+<]."
	cmds := Tokenize(testCode)
	for i, extepctedCmd := range []Cmd{
		{Op: INC_IND},
		{Op: RD_IN, Value: 28},
		{Op: CTRL_JMP, Value: 8},
		{Op: DEC_IND},
		{Op: INC_VAL},
		{Op: INC_IND},
		{Op: DEC_VAL},
		{Op: INC_IND},
		{Op: INC_VAL},
		{Op: DEC_IND},
		{Op: CTRL_RTN, Value: 8},
	} {
		assert.Equal(t, extepctedCmd, cmds[i])
	}
}

func TestTokenize_subjumps(t *testing.T) {
	testCode := ">[>[-<]<]>[-]+"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{Op: INC_IND},
		{Op: CTRL_JMP, Value: 7},
		{Op: INC_IND},
		{Op: CTRL_JMP, Value: 3},
		{Op: DEC_VAL}, {Op: DEC_IND},
		{Op: CTRL_RTN, Value: 3},
		{Op: DEC_IND},
		{Op: CTRL_RTN, Value: 7},
		{Op: INC_IND},
		{Op: CTRL_JMP, Value: 2},
		{Op: DEC_VAL},
		{Op: CTRL_RTN, Value: 2},
		{Op: INC_VAL},
	}
	require.Equal(t, len(cmds), len(expectedCmds))
	for i, extepctedCmd := range expectedCmds {
		assert.Equal(t, extepctedCmd, cmds[i], "mismatch on line %d", i)
	}
}

func TestTokenize_debug(t *testing.T) {
	testCode := ">#,28>>###"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{Op: INC_IND},
		{Op: WR_DEBUG, Value: 1},
		{Op: RD_IN, Value: 28},
		{Op: INC_IND},
		{Op: INC_IND},
		{Op: WR_DEBUG, Value: 3},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		assert.Equal(t, expectedCmd, cmds[i])
	}
}

func TestTokenize_inputs(t *testing.T) {
	testCode := ">,28>,>"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{Op: INC_IND},
		{Op: RD_IN, Value: 28},
		{Op: INC_IND},
		{Op: RD_IN, Value: -1},
		{Op: INC_IND},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		assert.Equal(t, expectedCmd, cmds[i])
	}
}

// TestTokenize_debugAndFlowControl This test was created to catch a bug with the debugging line messing up the
// stack pointer
func TestTokenize_debugAndFlowControl(t *testing.T) {
	testCode := "some text[>###[,28\ntesting sentence;\n>,>]]"
	cmds := Tokenize(testCode)
	expectedCmds := []Cmd{
		{Op: CTRL_JMP, Value: 9},
		{Op: INC_IND},
		{Op: WR_DEBUG, Value: 3},
		{Op: CTRL_JMP, Value: 5},
		{Op: RD_IN, Value: 28},
		{Op: INC_IND},
		{Op: RD_IN, Value: -1},
		{Op: INC_IND},
		{Op: CTRL_RTN, Value: 5},
		{Op: CTRL_RTN, Value: 9},
	}
	require.Equal(t, len(expectedCmds), len(cmds))
	for i, expectedCmd := range expectedCmds {
		require.Equal(t, expectedCmd, cmds[i], "error on command %d", i)
	}
}
