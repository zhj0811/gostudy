package main

import (
	// "github.com/hyperledger/fabric/common/metrics/prometheus"
	"github.com/gin-gonic/gin"
	// "github.com/prometheus/client_golang/prometheus"
	"github.com/peerfintech/BaseTools/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {

	// 运行模式
	gin.SetMode(gin.DebugMode) //ReleaseMode

	router := gin.Default()

	gauge := metrics.NewGauge(
		metrics.GaugeOpts{
			Namespace: "api",
			Name:      "celsius",
			Help:      "Current temperature of the CPU.",
		},
	)

	counter := metrics.NewCounter(
		metrics.CounterOpts{
			Namespace:  "api",
			Name:       "hd_errors_total",
			Help:       "Number of hard-disk errors.",
			LabelNames: []string{"device"},
		},
	)

	histogram := metrics.NewHistogram(
		metrics.HistogramOpts{
			Namespace:  "api",
			Name:       "test_histogram",
			Help:       "histogram test metrics",
			LabelNames: []string{"label1"},
		},
	)

	router.GET("/ping", func(c *gin.Context) {
		startTime := time.Now()
		counter.With("device", "/dev/sda").Add(1)
		c.JSON(200, gin.H{
			"message": "pong", //访问localhost:8080/ping页面显示{"message": "pong"}
		})
		histogram.With("label1", "test").Observe(time.Since(startTime).Seconds())
	})

	gauge.Set(1)

	// v1 := router.Group("/v1")
	// v1.GET("/login", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "v1 login")
	// })
	// v2 := router.Group("/v2")
	// v2.GET("/login", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "v2 login")
	// })
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9443", nil)
	}()
	router.Run(":8080") // listen and server on 0.0.0.0:8080(default)

	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":9090", nil)
}
