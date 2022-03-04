// @description 异常处理
//异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/2/22 3:07 下午
// @lastmodify 2022/2/22 3:07 下午

package main

import "fmt"

func Try (fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
func main()  {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
