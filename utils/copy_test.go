package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Testcopy_testSuite struct {
	suite.Suite
}

func (s *Testcopy_testSuite) TestFunc() {
	a := map[string]interface{}{
		"k": "v",
	}
	b := map[string]interface{}{}
	MapCopy(a, b)
	s.Equal(a["k"], b["k"])

	al := []map[string]interface{}{{"k1": "v1"}, {"k2": "v2"}}
	var bl []map[string]interface{}
	for _, value := range al {
		bl = append(bl, MapCopyValue(value))
	}
	s.Equal(bl[1]["k2"], al[1]["k2"])
	s.Equal(bl[0]["k1"], al[0]["k1"])
}

func (s *Testcopy_testSuite) SetupTest() {

}

func TestMapCopySuite(t *testing.T) {
	suite.Run(t, new(Testcopy_testSuite))
}
