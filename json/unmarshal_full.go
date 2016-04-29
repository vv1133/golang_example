package main

import "encoding/json"
import "fmt"

type Arr struct {
	Interger int    `json:"interger"` //首字母必须大写
	Sdata    string `json:"sdata"`
}

func main() {
	data := []byte(`{
                "single":{"a":"aaaa"},
                "array":[{"interger":1,"sdata":"abc"}, {"interger":2,"sdata":"ddd"}]
                }`)

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objmap); err != nil {
		panic(err)
	}

	fmt.Println(objmap)
	fmt.Println(*objmap["single"])
	fmt.Println(*objmap["array"])

	var single map[string]*json.RawMessage
	json.Unmarshal(*objmap["single"], &single)
	fmt.Println(single)

	var a string
	json.Unmarshal(*single["a"], &a)
	fmt.Println(a)

	/*
		var array []map[string]*json.RawMessage
		json.Unmarshal(*objmap["array"], &array)
		fmt.Println(array)
		fmt.Println(array[0])
		fmt.Printf("%T\n", array[0]["interger"])

		var aa int
		json.Unmarshal(*(array[0]["interger"]), &aa)
		fmt.Println(aa)

		var bb string
		json.Unmarshal(*(array[0]["sdata"]), &bb)
		fmt.Println(bb)
	*/
	var array []Arr
	fmt.Println("obj array", objmap["array"])
	json.Unmarshal(*objmap["array"], &array)
	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Printf("%T\n", array)
}
