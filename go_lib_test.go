package go_lib

import "fmt"
import "testing"

func init() {
	fmt.Println("Hello from go_lib!")
}

func TestMakeHello(t *testing.T) {
	name := "Denis"
	out := MakeHello(name)
	if out != "Hello, Denis!" {
		t.Error()
	}
}
