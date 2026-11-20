package go_lib

import "fmt"

func init() {
	fmt.Println("Hello from go_lib!")
}

func MakeHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
