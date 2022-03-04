// @description
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 6:41 PM
// @lastmodify 2022/3/3 6:41 PM

package main

import "fmt"

func closeChannel() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("结束")
}
