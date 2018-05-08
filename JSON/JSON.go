package main

import (
	"fmt"
	"math"
	"encoding/json"
	"os"
	"log"
)
//还未总结
func main() {
	//这里的interface {}是  任意的type，也就是generic type
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}
	//'{json格式}'表示为编码的json格式
	//我们一般说的encoding指的是将struct 编码为json格式
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	fmt.Println(b)
	var f interface{}
	json.Unmarshal(b, &f)
	//这个unmarshal后的是 map[string]interface{}类型的
	fmt.Println(f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}

	var m1 FamilyMember
	json.Unmarshal(b, &m1)
	fmt.Println(m1)

	// json.NewDecoder(os.Sdin) 创建一个object call dec
	//  var v map[string]interface{}   注意类型
	// dec.Decode(&v), 就把输入搞到了map v里面
	dec := json.NewDecoder(os.Stdin)
	fmt.Println(dec)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
