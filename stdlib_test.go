package stdlib

import "testing"

func TestMakeHello(t *testing.T) {
	name := "Denis"
	out := MakeHello(name)
	if out != "Hello, Denis!" {
		t.Error()
	}
}
