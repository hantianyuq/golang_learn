// @description
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 6:57 PM
// @lastmodify 2022/3/3 6:57 PM

package main

import "fmt"

func channelTest() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}
