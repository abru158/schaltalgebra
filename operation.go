package main

type operation func(l, r bool) bool

func opAND(l, r bool) bool {
	return l && r
}

func opOR(l, r bool) bool {
	return l || r
}
