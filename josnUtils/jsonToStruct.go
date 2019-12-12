package josnUtils

import (
	"fmt"
	"transfer/json"
)

type BaseJsonDemo struct {
	DataID int `json:"data_id"`
}

type JsonDemo struct {
	*BaseJsonDemo
	IP   string                   `json:"ip"`
	Topo []map[string]interface{} `json:"topo"`
}

func ToJsonDemo() {
	data := []byte(`{"ip":"10","data_id":1}`)
	res := &JsonDemo{}
	err := json.Unmarshal(data, res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.DataID)
}

func JsonToStructWithNil() {
	data := []byte(`{"ip":"10","data_id":null,"topo":null}`)
	res := &JsonDemo{}
	err := json.Unmarshal(data, res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.DataID)
}
