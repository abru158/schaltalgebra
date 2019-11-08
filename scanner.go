package main

type scanner struct {
	src []byte

	c   rune
	pos int
}

func (s *scanner) init(src []byte) {
	s.src = src
	s.pos = 0
	if len(src) > 0 {
		s.c = rune(src[0])
	}
}

func (s *scanner) next() {
	s.pos++
	if s.pos < len(s.src) {
		s.c = rune(s.src[s.pos])
	} else {
		s.c = -1
	}
}
