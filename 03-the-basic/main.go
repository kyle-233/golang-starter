package main

import (
	_ "fmt"
	ex "golang-starter/03-the-basic/example"
)

// **这里需要定义 go.mod 的module 为 golang-starte，否则这个不知道 golang-starter 是什么，导致找不到 example**

// 导入
// import "xxx"
// 别名
// import e "xxx"
// 批量导入

// import (
// 	"aaa"
//  "bbb"
// )

// 导入不调用
// import _ "fmt" // 如果该包下有 init 函数，会自动执行该函数

// 禁止循环导入

// 同一个包的不需要声明导入就可以使用对应的函数和变量

// 导出

// 导出通过命名区分的
// 1、首字母大写。该函数或者变量就可以对外暴露
// 2、首字母小写。私有，同一个包的可以引用

func main() {
	ex.SayHello()
}
