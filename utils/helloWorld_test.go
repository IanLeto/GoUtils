package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TesthelloWorld_testSuite struct {
	suite.Suite
}

func (s *TesthelloWorld_testSuite) TestFunc() {
	HelloWorld()
}

func (s *TesthelloWorld_testSuite) SetupTest() {

}

func TestHelloWorldSuite(t *testing.T) {
	suite.Run(t, new(TesthelloWorld_testSuite))
}
