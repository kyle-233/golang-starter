package main

// 下划线在导入中

// 表示不需要导入整个包，忽略其他的方法的导入，让其只执行 init 方法
// import _ "fmt"

// 下划线在代码中

// 表示忽略这个变量
// f, err := os.Open("xxx")
// f, _ := os.Open("xxx") // 忽略 err

func main() {

}
