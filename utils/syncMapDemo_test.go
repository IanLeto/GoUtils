package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestSyncMapDemoSuit struct {
	suite.Suite
}

func (s *TestSyncMapDemoSuit) SetupTest() {
}

func (s *TestSyncMapDemoSuit) TestUsageSyncMapDemo() {
	DemoSyncMap()
	time.Sleep(10 * time.Second)
}
func (s *TestSyncMapDemoSuit)TestUsageSyncMapDemo2()  {
	DemoSyncMap2()
}
func TestDemoSyncMapSuite(t *testing.T) {
	suite.Run(t, new(TestSyncMapDemoSuit))
}
