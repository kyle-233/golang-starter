package main

import "fmt"

// main 函数是go程序的默认入口函数（主函数）

func main() {
	fmt.Println("Main")
}

// init 函数用于包（package）的初始化
// 1、init 用于程序执行前做包的初始化
// 2、每个包可以拥有多个init 函数
// 3、包的每个源文件也可以拥有多个init函数
// 4、同一个包的init函数的执行顺序go语言没有明确的定义
// 5、不同包的init 函数按照包的导入的依赖关系决定初始化函数的执行顺序
// 6、init 函数不可以被调用，而是在main 函数前自动执行

func init() {
	fmt.Println("Init")
}

func init() {
	fmt.Println("Init2")
}

// 异同
// 相同：两个函数定义时不能有任何的参数和返回值，且Go程序自动调用
// 不同点：1、init函数可以应用于任意包中，且可以重复定义
// 2、main函数只可用于main包中，且只能定义一个

// 执行顺序
// 1、同一个go文件执行的顺序是从上到下
// 2、同一个package中的不同文件是按照文件名字符串比较 从小到大 的顺序调用各文件的init 函数
// 3、对于不同的package，如果不相互依赖的话，根据main包中的先引入的后调用的顺序执行，如果存在依赖，先调用最早被依赖的包的init，最后调用main函数
