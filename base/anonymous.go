// @description 
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/2/22 11:56 上午
// @lastmodify 2022/2/22 11:56 上午

package main

import (
	"fmt"
	"math"
)

func main() {
	// 匿名函数是指不需要定义函数名的一种函数实现方式， 在Go里面， 可以像普通变量一样被传递或使用，支持随时在代码里定义匿名函数。
	// 匿名函数由一个不带函数名的函数声明和函数体组成。 匿名函数的优越性在于可以在直接使用函数内的变量，不必申明。
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
	// ---- function variable
	fn := func(){
		println("Hello, World!")
	}
	fn()
	fus := []func(x int) int{
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}
	println(fus[0](100))
	println(fus[1](100))

	// --- function as field
	d := struct {
		fn func() string
	}{
		fn : func() string {
			return "Hello, World!"
	},
	}
	print(d.fn())

	// --- channel of function
	fc := make(chan func() string, 2)
	fc <- func() string {
		return "Hello, World!"
	}
	println((<-fc)())
}