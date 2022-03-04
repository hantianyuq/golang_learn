// @description
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/3/3 5:10 PM
// @lastmodify 2022/3/3 5:10 PM

package main

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
