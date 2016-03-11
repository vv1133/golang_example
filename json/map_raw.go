package main

import "encoding/json"
import "fmt"

func main() {
	data := []byte(`{"sendMsg":{"user":"ANisus","msg":"Trying to send a message"},"say":"Hello"}`)

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &objmap); err != nil {
		panic(err)
	}

	fmt.Println(objmap)
	fmt.Printf("%T\n", objmap)

	fmt.Println(*objmap["sendMsg"])
	fmt.Printf("%T\n", *objmap["sendMsg"])

	fmt.Println(*objmap["say"])
	fmt.Printf("%T\n", *objmap["say"])

	var str string
	json.Unmarshal(*objmap["say"], &str)
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	var sendMsg map[string]*json.RawMessage
	json.Unmarshal(*objmap["sendMsg"], &sendMsg)
	fmt.Println(sendMsg)
	fmt.Printf("%T\n", sendMsg)

	var user string
	json.Unmarshal(*sendMsg["user"], &user)
	fmt.Println(user)
	fmt.Printf("%T\n", user)
}
