// @description 延迟调用（defer） 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/2/22 3:07 下午
// @lastmodify 2022/2/22 3:07 下午

package main

/**
	defer特性
		1 关键字defer用于注册延迟调用。
		2 这些调用直到 return 前才被执行。 因此，可以用来做资源清理。
		3 多个defer语句， 按先进后出的方式执行。
		4 defer语句中的变量，在defer声明时就决定了。

	defer用途
		1 关闭文件句柄
		2 锁资源释放
		3 数据库连接释放
 */

func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y)
	}(x)

	x += 10
	y += 100
	println("x =", x, "y =", y)
}

func main()  {
	test()
}
