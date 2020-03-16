package prometheusDemo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func DocBaseDemo() {

}

func init() {
	var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temp",
		Help: "current temperature of cpu",
	})
	hdFailures := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "number of hard-disk",
	}, []string{"device"})

	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)

	cpuTemp.Set(88.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
