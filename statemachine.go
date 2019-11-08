package main

type stateMachine struct {
	s    map[rune]bool
	l, h rune
}

func (s *stateMachine) init(l, h rune) {
	s.s = map[rune]bool{}
	s.l, s.h = l, h
	// not required to have the below
	for i := l; i <= h; i++ {
		s.s[i] = false
	}
}

func (s *stateMachine) get(r rune) bool {
	return s.s[r]
}

func (s *stateMachine) increment() bool {
	for i, j := s.h, s.l; i >= j; i-- {
		cur := !s.s[i]
		s.s[i] = cur
		if cur {
			break
		}
		if i == j {
			return false
		}
	}
	return true
}
