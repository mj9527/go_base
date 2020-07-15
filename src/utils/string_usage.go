/**
 * @Author: mjzheng
 * @Description:
 * @File:  string_usage.go
 * @Version: 1.0.0
 * @Date: 2020/6/11 下午5:30
 */

package utils

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func IsSuffix(s string) bool {
	start := time.Now()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	end := time.Now()

	delta := end.Sub(start)
	fmt.Println(delta)
	return strings.HasSuffix(s, ".txt")
}
