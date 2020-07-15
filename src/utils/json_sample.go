/**
 * @Author: mjzheng
 * @Description:
 * @File:  json_sample.go
 * @Version: 1.0.0
 * @Date: 2020/7/14 下午7:14
 */

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	student "go_base/src/proto"
	"reflect"
	"strings"
)

type TPNSAcceptTime struct {
	Start  TPNSTimeElement   `json:"start"`
	End    TPNSTimeElement   `json:"end"`
	Ring   []TPNSTimeElement `json:"ring"`
	RoomId int64             `json:"room_id"`
}

type TPNSTimeElement struct {
	Hour string `json:"hour"`
	Min  string `json:"min"`
}

type TPNSAcceptTimeV1 struct {
	Start  TPNSTimeElement `json:"start"`
	Option int32           `json:"option"`
}

func StudyJson() {
	acceptTime := &TPNSAcceptTime{
		//Start: TPNSTimeElement{
		//	Hour: "10",
		//	Min: "59",
		//},
		End: TPNSTimeElement{
			Hour: "11",
			Min:  "01",
		},
		Ring: []TPNSTimeElement{
			{
				Hour: "1",
				Min:  "2",
			},
			{
				Hour: "3",
				Min:  "4",
			},
		},
		RoomId: 1000,
	}

	data, err := json.Marshal(acceptTime)
	if err != nil {
		fmt.Println("marshal error", err)
		return
	}

	fmt.Println("data", string(data))

	//newAccept := &TPNSAcceptTime{}
	//err = json.Unmarshal(data, newAccept)
	//if err != nil {
	//	fmt.Println("unmarshal error", err)
	//	return
	//}
	//
	//fmt.Printf("new %+v", newAccept)

	var u interface{}
	err = json.Unmarshal(data, &u)
	if err != nil {
		fmt.Println(err)
		return
	}

	ParseJson(u)
}

func ParseJson(u interface{}) {
	uMap, ok := u.(map[string]interface{})
	if !ok {
		fmt.Println("failed to get type")
		return
	}

	ParseMap(uMap)
}

func ParseMap(u map[string]interface{}) {
	//fmt.Println("map", u)
	for k, v := range u {
		//fmt.Println("map item", k, v)

		switch t := v.(type) {
		case string:
			fmt.Println(k, t)
		case []interface{}:
			fmt.Println(k, t)
			ParseArray(t)
		case map[string]interface{}:
			fmt.Println(k, t)
			ParseMap(t)
		case int64:
			fmt.Println(k, t)
		default:
			fmt.Printf("Type Square %T with value %v\n", t, t)
		}
	}
}

func ParseArray(t []interface{}) {
	//fmt.Println("array", t)
	for _, item := range t {
		//fmt.Println("array item", item)
		switch item := item.(type) {
		case string:
			fmt.Println(item)
		case map[string]interface{}:
			ParseMap(item)
		case []interface{}:
			ParseArray(item)
		}
	}
}

func StudyJsonPb() {
	jsonData := []byte(`{"name": "mjzheng", "age":18}`)

	handler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}

	st := student.Student{}

	fmt.Printf("%#v", st)
	return

	err := handler.Unmarshal(strings.NewReader(string(jsonData)), &st)
	if err != nil {
		fmt.Println("failed to parse", err)
		return
	}

	err = check(&st, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", st)

}

func check(i interface{}, s string) error {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := v.Kind()

	switch k {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			// 过滤非导出字段
			if !v.Field(i).CanInterface() {
				continue
			}
			// 过滤 pb XXX 默认字段
			if strings.Contains(t.Field(i).Name, "XXX") {
				continue
			}
			err := check(v.Field(i).Interface(), s+"."+t.Field(i).Name)
			if err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			return fmt.Errorf("gconf check nil, Name:%v, Kind:%v", s, t.Kind().String())
		}
		return check(v.Elem().Interface(), s)
	case reflect.Slice:
		if v.Len() == 0 {
			return fmt.Errorf("gconf check nil, Name:%v, Kind:%v", s, t.Kind().String())
		}
		for i := 0; i < v.Len(); i++ {
			err := check(v.Index(i).Interface(), s)
			if err != nil {
				return err
			}
		}
	default:
	}

	return nil
}
