package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	var message json.RawMessage
	var objmap map[string]*json.RawMessage

	message = []byte(`{"event": [{"event_name": "event0", "event_id": 1, "event_detail": "detail0"}, {"event_name": "event1", "event_id": 1, "event_detail": "detail1"}, {"event_name": "event2", "event_id": 2, "event_detail": "detail2"}], "ts_id": 0, "real_event_count": 3, "time_zone": 1.5, "service_id": 0}`)

	err := json.Unmarshal(message, &objmap)
	if err != nil {
		fmt.Println("error")
	}
	for k, v := range objmap {
		fmt.Printf("%s %s\n", k, v)
	}

	var keys []string
	for k := range objmap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s %s\n", k, objmap[k])
	}
}
