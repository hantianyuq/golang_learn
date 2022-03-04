// @description
// 这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。
// 让出CPU时间片，重新等待安排任务(大概意思就是本来计划的好好的周末出去烧烤，但是你妈让你去相亲,两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 5:47 PM
// @lastmodify 2022/3/3 5:47 PM

package main

import (
	"fmt"
	"runtime"
)

func goSched() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		// 切一下， 再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
