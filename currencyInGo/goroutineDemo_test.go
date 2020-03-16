package currencyInGo

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type TestgoroutineDemo_testSuite struct {
	suite.Suite
	server http.Server
}

func (s *TestgoroutineDemo_testSuite) TestFunc() {
	str := make(chan string, 0)
	DoWork(str)
	str <- "x"
	time.Sleep(4 * time.Second)
	close(str) // 如果不加上这一句，dowork 中的goroutine 将回泄漏
	time.Sleep(4 * time.Second)

}

func (s *TestgoroutineDemo_testSuite) SetupTest() {

}

func TestGoroutineDemoSuite(t *testing.T) {
	suite.Run(t, new(TestgoroutineDemo_testSuite))
}
