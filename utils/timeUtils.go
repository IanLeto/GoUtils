package utils

import (
	"time"
	"transfer/logging"
)

// 时间类型转纳秒时间戳转
func TimeDemo() {
	logger := logging.GetStdLogger()
	logger.Infof("时间类型 to 纳秒%v", time.Now().UnixNano())
	logger.Infof("时间类型 to 纳秒%v", time.Now().UnixNano())

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
	sec = tsec.Unix()
	msec = tmsec.UnixNano() / 1e6
	nsec = tnsec.UnixNano()

	// sec timestamp to msec timestamp
}

func TimeToNanoTimestamp(time time.Time) {

}
