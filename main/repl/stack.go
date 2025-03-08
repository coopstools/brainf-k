package repl

import "errors"

type stack struct {
	s []byte
	i int
}

func (s *stack) incInd() {
	s.i += 1
}

func (s *stack) decInd() error {
	s.i -= 1
	if s.i < 0 {
		return errors.New("index moved out of range")
	}
	return nil
}

func (s *stack) incVal() {
	s.s[s.i] += 1
}

func (s *stack) decVal() {
	s.s[s.i] -= 1
}

func (s *stack) val() byte {
	return s.s[s.i]
}

func (s *stack) out() byte {
	return s.s[s.i]
}

func (s *stack) in(i byte) {
	s.s[s.i] = i
}
