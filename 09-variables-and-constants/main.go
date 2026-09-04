package main

import "fmt"

// Error 短变量不可以用于全局的变量
// global := "global"
var global = "global"

func main() {
	// 标准声明
	// var 变量名 变量类型
	var name string
	var age int
	var isOk bool

	fmt.Println(name, age, isOk)

	// 批量声明
	var (
		name1 string
		age1  int
	)

	fmt.Println(name1, age1)

	// 初始化
	// var 变量名 变量类型 = 表达式
	var name2 string = "tom"
	var age2 int = 18

	// 一次初始化多个变量
	var name3, age3 = "tom", 18

	// 类型推导（可以省略类型声明）
	var name4 = "peter"

	fmt.Println(name2, age2, name3, age3, name4)

	// 短变量声明
	name5 := "marry"

	fmt.Println(global, name5)

	// 匿名变量
	// 不占用命名空间，不占用内存，不存在重复声明
	name6, _ := foo()

	fmt.Println(name6)

	// 常量
	// 声明和变量声明一样，只是使用const
	const url = "www.xxx.com"
	const (
		post = 3000
		time = ""
	)

	// iota
	// 只可以跟const一起使用
	const (
		a = iota
		b
		c
	)
}

func foo() (string, int) {
	return "tom", 18
}
