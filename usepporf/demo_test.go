package usepporf

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Testdemo_testSuite struct {
	suite.Suite
}

func (s *Testdemo_testSuite) TestFunc() {
	WorkerManager(10)
}

func (s *Testdemo_testSuite) SetupTest() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Testdemo_testSuite))
}
