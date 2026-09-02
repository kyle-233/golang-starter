package example

import "fmt"

func sayHi() {
	fmt.Println("say hi")
	// SayHello() // 这里会导致循环引用
}

func Say() {
	fmt.Println("Say")
}
