// @description 延迟调用（defer）*滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/2/22 3:07 下午
// @lastmodify 2022/2/22 3:07 下午

package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func test1() {
	lock.Lock()
	lock.Unlock()
}

func test1defer() {
	lock.Lock()
	defer lock.Unlock()
}

func main()  {
	func() {
		t1 := time.Now()
		
		for i :=0; i < 10000; i ++ {
			test1()
		}
		
		elapsed := time.Since(t1)
		fmt.Println("test elapsed:", elapsed)
	}()
	
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i ++ {
			test1defer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testder elapsed:", elapsed)
	}()
}
