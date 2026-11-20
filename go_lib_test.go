package go_lib

import "testing"

func TestMakeHello(t *testing.T) {
	name := "Denis"
	out := MakeHello(name)
	if out != "Hello, Denis!" {
		t.Error()
	}
}
