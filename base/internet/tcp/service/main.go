// @description 
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/1 5:21 下午
// @lastmodify 2022/3/1 5:21 下午

// TCP协议
// TCP/IP即传输控制协议/网间协议，是一种面向连接的、可靠的、基于字节流的传输层通信协议，因为是面向链接的协议，
//  数据像水流一样传输，会存在黏包问题。

// 一个TCP服务端可以同时连接很多歌客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。因为GO语言中创建多个
// goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。
// TCP服务端程序的处理流程：
//  1.监听端口 2.接受客户端请求建立链接 3.创建goroutine处理链接


package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭链接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))//发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for  {
		fmt.Println("listen", listen)
		conn, err := listen.Accept() // 建立链接
		fmt.Println("Accept", conn)
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}