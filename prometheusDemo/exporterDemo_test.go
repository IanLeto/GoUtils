package prometheusDemo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestexporterDemo_testSuite struct {
	suite.Suite
}

func (s *TestexporterDemo_testSuite) TestFunc() {
	DocBaseDemo()
}

func (s *TestexporterDemo_testSuite) SetupTest() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestexporterDemo_testSuite))
}
