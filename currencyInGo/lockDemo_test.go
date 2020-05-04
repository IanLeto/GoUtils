package currencyInGo

import (
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"transfer/logging"
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
	var x sync.Once

	for _, value := range []int{1, 2, 3, 45, 12} {
		logging.GetStdLogger().Info("xx", value)
		if value == 3 {
			x.Do(func() {
				logging.GetStdLogger().Warn("xcxccx", value)
			})
		}

	}

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestlockDemo_testSuite))
}
