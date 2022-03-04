// @description
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 3:55 PM
// @lastmodify 2022/3/3 3:55 PM

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
