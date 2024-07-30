package hw_test

import (
	"bytes"
	"coopstools/brainf-k/main/hw"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestHW1(t *testing.T) {
	for k, f := range map[string]func(writer io.Writer){"hw1": hw.HW1, "hw2": hw.HW2, "hw3": hw.HW3} {
		t.Run(k, func(t *testing.T) {
			buf := bytes.Buffer{}
			f(&buf)
			assert.Equal(t, "HELLO, WORLD!!", buf.String())
		})
	}
}
