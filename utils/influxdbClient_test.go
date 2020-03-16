package utils

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestinfluxdbClient_testSuite struct {
	suite.Suite
	cli client.Client
}

func (s *TestinfluxdbClient_testSuite) TestFunc() {

}

func (s *TestinfluxdbClient_testSuite) SetupTest() {
	//s.cli, err := NewInfluxDBClinet()
	var err error
	s.cli, err = NewInfluxDBClinet()
	s.NoError(err)
	_, _, err = s.cli.Ping(1000)
	s.NoError(err)
}

func (s *TestinfluxdbClient_testSuite) TestCreateTable() {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "ns",
	})
	point, err := client.NewPoint("t1", map[string]string{"k": "v"}, map[string]interface{}{"k": "v"}, time.Unix(0, T1()))
	points.AddPoint(point)
	s.NoError(err)
	s.NoError(s.cli.Write(points))
}

// s pre
func T1() int64 {
	return time.Now().Unix()
}

func T2() int64 {
	return time.Now().UnixNano()

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestinfluxdbClient_testSuite))
}
