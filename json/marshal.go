package main

import (
	"fmt"
        "encoding/json"
)

type EpgAttrs struct {
        EleInt int
        EleStr string
        EleSlice []string
}

func main() {
        epg_attrs := &EpgAttrs{
                EleInt: 1,
                EleStr: "hello",
                EleSlice: []string{"apple", "pear"},
        }
        attr, _ := json.Marshal(epg_attrs)
        fmt.Println(string(attr))
}

