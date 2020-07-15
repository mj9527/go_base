/**
 * @Author: mjzheng
 * @Description:
 * @File:  time_sample.go
 * @Version: 1.0.0
 * @Date: 2020/7/15 下午4:32
 */

package utils

import (
	"fmt"
	"time"
)

func TimeSample() {
	t := time.Now()
	fmt.Printf("%d.%d.%d\n", t.Year(), t.Month(), t.Day())
	fmt.Println(t)
	fmt.Println(t.UTC())
	fmt.Println(t.Unix())
	fmt.Println(t.UnixNano())
	fmt.Println(t.Format(time.RFC1123))
	defer func() {
		fmt.Println(time.Now().Sub(t))
		fmt.Println(time.Since(t))
	}()
}
