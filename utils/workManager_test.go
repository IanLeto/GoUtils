package utils

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"transfer/logging"
)

type TestworkManager_testSuite struct {
	suite.Suite
}

func Dosomething(a interface{}) {

}
func (s *TestworkManager_testSuite) TestFunc() {
	x := make(chan int)
	go func() {
		defer fmt.Println("我好了")
		for value := range x {
			logging.GetStdLogger().Info(value)

		}
	}()
	x <- 1
	time.Sleep(5 * time.Second)
	logging.GetStdLogger().Info(2)
	close(x)

	logging.GetStdLogger().Info(3)
	close(x)

}

func (s *TestworkManager_testSuite) SetupTest() {

}

func TestWorkSuite(t *testing.T) {
	suite.Run(t, new(TestworkManager_testSuite))
}
