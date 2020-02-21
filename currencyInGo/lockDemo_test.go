package currencyInGo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestlockDemo_testSuite struct {
	suite.Suite
}

func (s *TestlockDemo_testSuite) TestFunc() {}
func (s *TestlockDemo_testSuite) TestDemoOnce() {
	DemoOnce()
	DemoOnce2()
}
func (s *TestlockDemo_testSuite) TestDemo2() {
	Demo2()
	DemoOnce2()
}

func (s *TestlockDemo_testSuite) TestDemoLock() {
	DemoLock()
	DemoWithoutLock()
}

func (s *TestlockDemo_testSuite) SetupTest() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestlockDemo_testSuite))
}
