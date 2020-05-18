package utils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"hash/fnv"
	"testing"
	"time"
)

type TesthashDemo_testSuite struct {
	suite.Suite
}

func (s *TesthashDemo_testSuite) TestFunc() {

}

func (s *TesthashDemo_testSuite) SetupTest() {
	hash := fnv.New32a()

	//var x = fmt.Sprintf("%d", hash.Sum32())
	x, err := hash.Write([]byte("10.0.1.68:10086"))
	if err != nil {
		panic(err)
	}
	fmt.Println(hash.Sum32())
	fmt.Println(x)
	ss := time.Duration(30000000000)
	fmt.Println(ss.Seconds())
}

func TestHashDemoSuite(t *testing.T) {
	suite.Run(t, new(TesthashDemo_testSuite))
}
