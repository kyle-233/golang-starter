package main

import "fmt"

func main(){
	a := 10
	b := &a

	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)

	// 指针取值
	fmt.Printf("type of b:%T\n", b)

	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

	// 空指针
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s\n", p)

	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

	// 报错例子 panic
	// panic: runtime error: invalid memory address or nil pointer dereference
	// var aPoint *int
	// *aPoint = 100
	// fmt.Println(*aPoint)

	// panic: assignment to entry in nil map
	// var bMap map[string]int
	// bMap["测试"] = 100
	// fmt.Println(bMap)

	// 为什么报错？
	// 因为引用类型的变量在声明时需要初始化内存大小（分配内存空间），值类型的变量不需要，因为在声明时会自动分配内存空间
	// 需要分配内存空间时使用 new 和 make

	// new
	aNew := new(int)
	bNew := new(bool)
	fmt.Printf("%T\n", aNew)
	fmt.Printf("%T\n", bNew)
	fmt.Println(*aNew)
	fmt.Println(*bNew)

	var aPoint_ *int // 声明一个指针变量，但是没有初始化，指针是一个引用类型
	aPoint_ = new(int) // 使用new初始化，分配内存空间
	*aPoint_ = 10 // 取值
	fmt.Println(*aPoint_)

	// 对于 slice map, chan需要使用 make 初始化内存
	var bMap_ map[string]int
	bMap_ = make(map[string]int, 10)
	bMap_["测试"] = 100
	fmt.Println(bMap_)

	var num int = 99
	fmt.Println("num:", num)
	fmt.Printf("num address:%p\n", &num)
	var numPtr *int
	numPtr = &num
	fmt.Println("numPtr:", numPtr)
	*numPtr = 100
	fmt.Println("num:", num)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}