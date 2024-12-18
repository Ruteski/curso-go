﻿package main

import (
	"fmt"
	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser

	jsonData := `{"foo": "bar", "num": 123, "bool": true, "arr": [1,2,3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("num=%d\n", v.GetInt("num"))
	fmt.Printf("bool=%t\n", v.GetBool("bool"))
	fmt.Printf("arr=%v\n", v.GetArray("arr"))

	a := v.GetArray("arr")
	for i, value := range a {
		fmt.Printf("index %d: %s\n", i, value.String())
		fmt.Printf("index=%d\n", i)
		fmt.Printf("val=%s\n", value)
	}
}
