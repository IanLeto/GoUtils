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
	//s.timeDemo = time.Now()
	//logger := logging.GetStdLogger()
	//
	//// int64 to timestamp
	//logger.Info("golang 没有所谓的time.timestamp类型，int 就是时间戳，这里与python不同")
	//var sec int64 = 1582701190
	//var msec int64 = 1582701190946
	//var nsec int64 = 1582701190946414000
	//// timestamp to time
	//tsec := time.Unix(sec, 0)
	//tmsec := time.Unix(0, msec*int64(time.Millisecond))
	//tnsec := time.Unix(0, nsec)
	//
	//// time to timestamp
	//s.Equal(sec, tsec.Unix())
	//s.Equal(msec, tmsec.UnixNano()/1e6)
	//s.Equal(nsec, tnsec.UnixNano())
	//
	//// sec timestamp to msec timestamp
	//logger.Info(sec * int64(time.Millisecond))
	//logger.Info(sec * int64(time.Nanosecond))
	//xx := time.Duration(sec)
	//fmt.Println("cx", xx)
	//
	//// time to 标准时间格式
	//logger.Info(tsec.Format(time.RFC3339))
	//logger.Info(tnsec.Format(time.RFC3339Nano))
	//// 解析字符串时间
	//hours, _ := time.ParseDuration("10h")
	//complex, _ := time.ParseDuration("1h10m10s")
	//logger.Infof("10h => %v", hours)
	//logger.Infof("second: %v", complex.Seconds())
	//
	//logger.Info("-----------")
	//a := time.Unix(1583311518000000000, 0)
	//b := time.Unix(0, 1583311518000000000)
	//
	//logger.Infof("%v", a.UnixNano())
	//logger.Infof("%v", a.Unix())
	//
	//logger.Infof("%v", b.UnixNano())
	//logger.Infof("%v", b.UnixNano())
	//
	//c := time.Unix(0, 1582701190)
	//logger.Info(c.Unix())
	//logger.Info(c.UnixNano())
	rr(1582701190)
	//rr(1583311518000000000)
	//
	//logger.Info(time.Now().UnixNano())
	//logger.Info(time.Unix(0,1583982851473426).Format(time.RFC3339Nano))
	//logger.Info(time.Unix(0,1583983973246439000).Format(time.RFC3339Nano))
	//logger.Info(time.Unix(0,1583983973246439).Format(time.RFC3339))
}

func rr(n int64) () {
	logger := logging.GetStdLogger()
	logger.Infof("xc: %v", time.Unix(n, 0).Unix())
	logger.Infof("xc: %v", time.Unix(n, 0).UnixNano())

	logger.Infof("xc: %v", time.Unix(0, n).Unix())

	logger.Infof("xc:d %v", time.Unix(0, n).UnixNano())

}

func TestTimeDemoSuite(t *testing.T) {
	suite.Run(t, new(TesttimeUtils_testSuite))
}
