package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"transfer/logging"
)

type TesttimeUtils_testSuite struct {
	suite.Suite
	timeDemo time.Time
}

func (s *TesttimeUtils_testSuite) TestFunc() {
	//var logger = logging.GetStdLogger()
	//stamp := time.Unix(1531293019, 0)
	//logger.Infof("时间戳转时间类型：%v ", stamp)
	//logger.Infof("时间类型 ->纳秒时间戳：%v", s.timeDemo.UnixNano())
	//millisecond := s.timeDemo.UnixNano() / 1e6
	//logger.Infof("时间类型 ->毫秒时间戳：%v", millisecond)
	//logger.Infof("毫秒时间戳 -> 时间类型 %v", time.Unix(0, millisecond*int64(time.Millisecond)))
	//logger.Infof("纳秒时间戳 -> 时间类型 %v", time.Unix(0, s.timeDemo.UnixNano()))
	//logger.Info(s.timeDemo.Unix())

}

func (s *TesttimeUtils_testSuite) SetupTest() {
	s.timeDemo = time.Now()
	logger := logging.GetStdLogger()

	// int64 to timestamp
	logger.Info("golang 没有所谓的time.timestamp类型，int 就是时间戳，这里与python不同")
	var sec int64 = 1582701190
	var msec int64 = 1582701190946
	var nsec int64 = 1582701190946414000
	// timestamp to time
	tsec := time.Unix(sec, 0)
	tmsec := time.Unix(0, msec*int64(time.Millisecond))
	tnsec := time.Unix(0, nsec)

	// time to timestamp
	s.Equal(sec, tsec.Unix())
	s.Equal(msec, tmsec.UnixNano()/1e6)
	s.Equal(nsec, tnsec.UnixNano())

	// sec timestamp to msec timestamp
	logger.Info(sec * int64(time.Millisecond))

	// time to 标准时间格式
	logger.Info(tsec.Format(time.RFC3339))
	logger.Info(tnsec.Format(time.RFC3339Nano))

}

func TestTimeDemoSuite(t *testing.T) {
	suite.Run(t, new(TesttimeUtils_testSuite))
}
