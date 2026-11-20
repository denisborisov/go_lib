package go_lib

import "fmt"

func init() {
	fmt.Println("Hello from go_lib v2!")
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
