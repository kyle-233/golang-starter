# 内置类型

## 值类型
- bool
- int32 int64 int8 int16
- uint32 uint64 uint8(byte) uint16
- float32 float64
- string
- comples64 complex128
- array -固定长度的数组

## 引用类型
- slice -序列数组
- map -映射
- chan -管道

# 内置函数
Go 拥有一个不需要导入即可直接使用的函数
- append -用来追加元素到数组、slice,返回修改后的数组、slice
- close -主要用来关闭channel
- delete -从map中删除key对应的value
- panic -停止常规的goroutine (panic 和 recover：用来做错误处理)
- recover -允许程序定义goroutine的panic动作
- imag -返回complex的实部 （complex、real imag: 用于创建和操作基数 ）
- real -返回complex的虚部
- make -用来分配内存，返回Type本省（只能应用于slice, map, channel）
- new -用来分配内存，主要用来分配值类型。比如int，struct。返回指向Type的指针
- cap -capacity 容量的意思，用于返回某个类型的最大容量（只能用于切片和map）
- copy -用于复制和连接slice，返回复制的数目
- len -求长度，用于返回长度，string array slice map channel