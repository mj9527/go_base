/**
 * @Author: mjzheng
 * @Description:
 * @File:  string_sample.go
 * @Version: 1.0.0
 * @Date: 2020/7/15 下午5:23
 */

package utils

import (
	"bytes"
	"fmt"
	"strings"
)

func StringSample() {
	//s:= "a=1;b=2;c=3"
	//pairs := strings.Split(s, ";")
	//param := strings.Join(pairs, "&")
	//fmt.Println(param)
	//for _, pair := range pairs {
	//	words := strings.Split(pair, "=")
	//
	//	fmt.Println(words)
	//}
	//
	s1 := "cut  my   cut    name  cut"
	s2 := strings.Replace(s1, "cut", "", -1)
	s3 := strings.TrimSpace(s2)
	fmt.Println(s3)
	//data := make ([]byte, len(s3))
	//copy(data, s3)

	data := []byte(s3)
	var b []byte
	isFirst := false
	for _, c := range data {
		if c == ' ' {
			if isFirst {
				continue
			}
			isFirst = true
		} else if c != ' ' {
			isFirst = false
		}
		b = append(b, c)
	}

	fmt.Println(string(b))

	r1 := bytes.NewReader([]byte("my name is zheng"))
	buf := make([]byte, 256)
	n, err := r1.Read(buf)
	if err != nil {
	}
	fmt.Println(n, string(buf))

	//r2 := strings.NewReader(string("my name is zheng jun ming"))

}
