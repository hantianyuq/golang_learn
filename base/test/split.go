// @description 
// @author hantianyu
// Copyright 2022 sndks.com. All rights reserved.
// @datetime 2022/2/22 5:02 下午
// @lastmodify 2022/2/22 5:02 下午

package test

import (
	"strings"
)

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}