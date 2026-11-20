package stdlib

import "fmt"

func MakeHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
