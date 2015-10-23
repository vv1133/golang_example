package main

import "encoding/json"
import "fmt"

type Head struct {
	Head1 int
	Head2 string
}

type TestField1 struct {
	FieldStr string `json:"fieldstr"`
}

type TestField2 struct {
	FieldInt int `json:"fieldint"`
}

type Response struct {
	Head
	Page      int             `json:"page"`
	TestField json.RawMessage `json:"testfield"`
}

func main() {
	str1 := `{"Head1":11, "Head2":"headdddd", "page": 1, "testfield": {"fieldstr": "ffffffffffffff"}}`
	str2 := `{"Head1":11, "Head2":"headdddd", "page": 2, "testfield": {"fieldint":99}}`
	res := Response{}

	f := func(res *Response) {
		if res.Page == 1 {
			ttt := TestField1{}
			json.Unmarshal(res.TestField, &ttt)
			fmt.Println(ttt)
		} else if res.Page == 2 {
			ttt := TestField2{}
			json.Unmarshal(res.TestField, &ttt)
			fmt.Println(ttt)
		}
	}

	json.Unmarshal([]byte(str1), &res)
	fmt.Println(res)
	f(&res)

	json.Unmarshal([]byte(str2), &res)
	fmt.Println(res)
	f(&res)
}
