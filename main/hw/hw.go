package hw

import (
	"fmt"
	"io"
)

func HW1(w io.Writer) {
	var q1 uint = 0x01B2F02FB28
	var q2 uint = 0x01932DCCB25
	for i := 0; i < 7; i++ {
		for _, q := range []*uint{&q1, &q2} {
			r := (*q & 0x3F) + 0x20
			_, _ = w.Write([]byte{byte(r)})
			*q >>= 6
		}
	}
}

func HW2(w io.Writer) {
	chars := " ,!DEHLORW"
	msg := 0x22368790176645
	var char uint8
	for i := 0; i < 14; i++ {
		char = chars[msg&0xF]
		_, _ = fmt.Fprintf(w, "%c", char)
		msg >>= 4
	}
}

func HW3(w io.Writer) {
	var A, B, C, D, nA, nB, nC, nD byte
	var mask, v byte = 0b1, 0b0

	for ii := 0; ii < 14; ii++ {
		v = 0b0
		i := ii
		for _, l := range []*byte{&D, &C, &B, &A} {
			*l = byte(i) & mask
			i >>= 1
		}
		nA, nB, nC, nD = A^mask, B^mask, C^mask, D^mask
		s0 := B & C & D
		s1 := nA & B & nC

		v += nB | (C & D) | (nA & nC & nD) // MSB
		v <<= 1
		v += (A & nB & nC & D) | s0
		v <<= 1
		v += (nB & nD) | (nA & nB & C) | s1
		v <<= 1
		v += (nA & D) | (nB & C) | s1 | (A & nB & nD)
		v <<= 1
		v += (A & nB & nC) | s0 | (nA & B & nC & nD)
		v <<= 1
		v += (A & B) | (B & nC & nD) | (A & nC & nD) | s0 | (nA & nB & nC & D) //LSB
		_, _ = fmt.Fprintf(w, "%c", v+32)
	}
}
