package rennbon

import "testing"

type ucTest struct {
	in, out string
}

var test = &ucTest{"abc", "abc你好"}

func TestUC(t *testing.T) {
	a := SayHello(test.in)
	if a != test.out {
		t.Error()
	}
}
