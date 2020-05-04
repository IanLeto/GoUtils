package regexDemo_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"regexp"
	"strings"
	"testing"
)

type TestconsulRe_testSuite struct {
	suite.Suite
}

func (s *TestconsulRe_testSuite) TestFunc() {
	str := "bk_bkmonitorv3_enterprise_production/metadata/data_id/manual/bkmonitorv3-2068018684/1500419"
	ok, err := regexp.MatchString("data_id/manual", str)
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println(str)
		st := strings.Split(str, "/")
		fmt.Println(st[len(st)-1])
		fmt.Println(st[len(st)-2])

	}

}

func (s *TestconsulRe_testSuite) SetupTest() {

}

func TestConsulReSuite(t *testing.T) {
	suite.Run(t, new(TestconsulRe_testSuite))
}
