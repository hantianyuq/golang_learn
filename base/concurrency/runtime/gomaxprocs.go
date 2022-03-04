// @description
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 6:17 PM
// @lastmodify 2022/3/3 6:17 PM

package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}
func goMaxProCs() {
	runtime.GOMAXPROCS(5)
	go a()
	go b()
	time.Sleep(time.Second)
}
