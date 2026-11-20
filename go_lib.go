package go_lib

import "fmt"

func init() {
	fmt.Println("Hello from go_lib!")
}

func MakeHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func Double(num int) int {
	return num * 2
}


func Triple(num int) int {
	return num * 3
}
