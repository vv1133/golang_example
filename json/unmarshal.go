package main

import "encoding/json"
import "fmt"

type Head struct {
	Head1 int
	Head2 string
}

type TestField struct {
	Field1 int
	Field2 string
}

type Response2 struct {
	Head
	Page      int      `json:"page"`
	Fruits    []string `json:"fruits"`
	TestField `json:"field"`
}

func main() {
	str := `{"Head1":11, "Head2":"headdddd", "page": 1, "fruits": ["apple", "peach"], "field": {"Field1":3,"Field2":"heo"}}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)

	//res := Response2{Head{222,"aaaaa"}, 3,[]string{"hi","ui"},TestField{34,"hello"}}
	//attr, _ := json.Marshal(res)
	//fmt.Println(string(attr))

	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
