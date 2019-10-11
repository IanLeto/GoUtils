package Counterexample_test

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"testing"
	"transfer/json"
)

var path = "E:\\configCollect"

func TestGetNewRedisConfig(t *testing.T) {
	//conf := GetNewRedisConfig()
	res, _ := json.Marshal(nil)

	fmt.Println(conv.String(res))
}
