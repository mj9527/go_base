/**
 * @Author: mjzheng
 * @Description:
 * @File:  home_score
 * @Version: 1.0.0
 * @Date: 2020/6/18 下午7:50
 */

package utils

import (
	"fmt"
	"unsafe"
)

type HomeScore struct {
	Uid   int `json:"uid" xorm:"not null pk BIGINT(20)"`
	Score int `json:"score" xorm:"FLOAT(64)"`
}

type HomeScoreS []*HomeScore

var publicLs HomeScoreS

func GetHomeAnchorScoreList() {
	//var ls HomeScoreS
	ls := make(HomeScoreS, 10)
	for i := 0; i < cap(ls); i++ {
		//fmt.Println(len(ls), cap(ls))
		score := &HomeScore{
			Uid:   i,
			Score: i,
		}
		ls[i] = score
		fmt.Println(ls[i])
		//ls =append(ls, score)
	}

	publicLs = ls
	fmt.Println(publicLs)
	for i, item := range publicLs {
		fmt.Println(&publicLs[i], publicLs[i], item)
	}

	GetString()
}

func GetString() {
	s := "my name is my zheng"

	data := []byte(s)
	data[0] = 'h'

	fmt.Println(s)
	fmt.Println(data)

	p := new([5]int)
	fmt.Println("pointer len", unsafe.Sizeof(p))
}
