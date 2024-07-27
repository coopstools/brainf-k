package runner

import "errors"

type stack struct {
	s []byte
	i int
	c int64
}

func (s *stack) incInd() {
	s.c += 1
	s.i += 1
	//fmt.Println("inc i", s.i, s.s[:2])
}

func (s *stack) decInd() error {
	s.c += 1
	s.i -= 1
	if s.i < 0 {
		return errors.New("index moved out of range")
	}
	return nil
	//fmt.Println("dec i", s.i, s.s[:2])
}

func (s *stack) incVal() {
	s.c += 1
	s.s[s.i] += 1
}

func (s *stack) decVal() {
	s.c += 1
	s.s[s.i] -= 1
	//fmt.Println("dec v", s.i, s.s[:2])
}

func (s *stack) val() byte {
	s.c += 1
	return s.s[s.i]
}

func (s *stack) out() byte {
	s.c += 1
	return s.s[s.i]
}

func (s *stack) in(i byte) {
	s.c += 1
	s.s[s.i] = i
}
