package utils

import "github.com/influxdata/influxdb/client/v2"

func NewInfluxDBClinet() (client.Client, error) {
	conf := client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
	}
	return client.NewHTTPClient(conf)
}
