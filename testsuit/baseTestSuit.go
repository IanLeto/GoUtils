package testsuit

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// testify 提供了一个很优秀的类型断言
func TestCalc(t *testing.T) {
	assert.Equal(t, 1, 1)
}

// 我们现在可以为测试套件来设置生命周期

// 可以理解为一个固定用法
type CliSuite struct {
	suite.Suite
}

func (s *CliSuite) SetupTest() {

}

// 此时，我们可以用提供方法的方式来测试
func (s *CliSuite) TestCalc() {
	s.Equal(1, 1)
}

// 然而，我们需要一个类似于init方法来让整个test suit（测试套件）运行起来
func TestSuite(t *testing.T) {
	suite.Run(t, new(CliSuite))
}
