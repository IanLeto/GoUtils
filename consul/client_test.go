package consul

import (
	"GoUtils/utils"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"time"
	"transfer/logging"
)

type TestConsulClientSuite struct {
	suite.Suite
	Cli *ConfigCenter
}

func (s *TestConsulClientSuite) SetupTest() {
	s.Cli = NewConsulConfig(nil)
}

func (s *TestConsulClientSuite) TestKVGet() {
	value, _, err := s.Cli.KV.Get("ian/", nil)
	s.NoError(err)
	fmt.Println(value)
	//s.Equal("hello consul", string(value.Value))
	rs := []byte("xc")

	_, err = s.Cli.KV.Put(&api.KVPair{
		Key:   "ian/hello/world",
		Value: rs,
	}, nil)
	s.NoError(err)
}

func (s *TestConsulClientSuite) TestKVPut() {

}

func (s *TestConsulClientSuite) TestWatch() {
	watchConfig := make(map[string]interface{})

	//watchConfig["type"] = "keyprefix"
	watchConfig["prefix"] = "ian/"
	watchPlan, err := watch.Parse(watchConfig)

	if err != nil {
		panic(err)
	}
	watchPlan.Handler = func(u uint64, i interface{}) {
		services := i.(api.KVPairs)
		str, err := json.Marshal(services)
		utils.NoErr(err)
		fmt.Println(string(str))
	}
	if err := watchPlan.Run("localhost:8500"); err != nil {
		panic(err)
	}
	time.Sleep(1000 * time.Second)

}
func (s *TestConsulClientSuite) TestKVInTransfer() {
	//v, _, err := s.Cli.KV.Get("bkmonitor/service/1001", nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(conv.String(v.Value))

}

func (s *TestConsulClientSuite) TestKVKeys() {
	v, _, err := s.Cli.KV.Keys("bkmonitor/service", "1", nil)
	if err != nil {
		panic(err)
	}
	for _, value := range v {
		fmt.Println(value)
	}

}

func BenchmarkExtractdepart(b *testing.B) {
	cases := []map[string]string{{
		"t1": "t1V",
		"t2": "t2V",
	}, {"t1": "t1V",
		"t2": "t2V",},
		{"t1": "t1V",
			"t2": "t2V",}}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//departeq(cases[0], cases[1])
		RemoveRepByLoopDepart(cases)
	}
	b.StopTimer()

}

func BenchmarkExtractReflect(b *testing.B) {
	cases := []map[string]string{{
		"t1": "t1V",
		"t2": "t2V",
	}, {"t1": "t1V",
		"t2": "t2V",},
		{"t1": "t1V",
			"t2": "t2V",}}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//departeq(cases[0], cases[1])
		RemoveRepByLoopReflect(cases)
	}
	b.StopTimer()

}

func RemoveRepByLoopReflect(slc []map[string]string) []map[string]string {
	var result []map[string]string
	{
	} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if reflect.DeepEqual(slc[i], result[j]) {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func RemoveRepByLoopDepart(slc []map[string]string) []map[string]string {
	var result []map[string]string
	{
	} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if departeq(slc[i], result[j]) {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func departeq(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}

func (s *TestConsulClientSuite) TestHealth() {
	a, _, _ := s.Cli.Client.Health().Service("httpServer", "", false, nil)
	logger := logging.GetStdLogger()
	for _, value := range a {
		logger.Infof("value.Service.ID %v", value.Service.ID)
		logger.Infof("value.Service.Service %v", value.Service.Service)
		logger.Infof("value.Service.Tags %v", value.Service.Tags)
	}
	//health := s.Cli.Client.Health()
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(TestConsulClientSuite))
}
