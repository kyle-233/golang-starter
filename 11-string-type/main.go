package main

import "fmt"


func main() {
	// 字符串声明： "" / ``
	s1:= "您好"
	fmt.Println(s1)

	// 转义符：
	// \r -回车符
	// \n -换行符
	// \t -制表符
	// \' -单引号
	// \" -双引号
	// \ -反斜杠
	fmt.Println("str := \"c:\\pprof\\main.go\"")

	// 多行字符串 使用 ``
	s2 := `
	第一行
	第二行
	第三行
	`
	fmt.Println(s2)

	// 字符串的常用方法
	// 求长度 -len(str)
	// 拼接字符串 - + 或者 fmt.SPrintf()
	// 分割字符串 -strings.Split
	// 判断是否包含 -strings.Contains()
	// 前缀/后缀判断 -strings.HasPrefix, strings.HasSuffix
	// 子串出现的位置 -strings.Index(), strings.LastIndex()
	// join操作 -strings.Join(a[]string, sep string)

	// 遍历字符串
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()

	// 字符串底层是 byte 数组，可以与 []byte 互转

	// 修改字符串
	// 字符串不可以直接修改，需要转成 []rune 或者 []byte，然后再转成字符串
	s3 := "hello"
	byteS3 := []byte(s3)
	byteS3[0] = 'H'
	fmt.Println(string(byteS3))

	s4 := "博客"
	runeS4 := []rune(s4)
	runeS4[0] = '狗'
	fmt.Println(string(runeS4))
}