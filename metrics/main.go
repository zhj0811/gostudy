package main

import (
	// "github.com/hyperledger/fabric/common/metrics/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {

	// 运行模式
	gin.SetMode(gin.DebugMode) //ReleaseMode

	router := gin.Default()

	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})

	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)

	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "test_histogram",
			Help: "histogram test metrics",
		},
		[]string{"label1"},
	)

	prometheus.MustRegister(gauge)
	prometheus.MustRegister(counter)
	prometheus.MustRegister(histogram)

	router.GET("/ping", func(c *gin.Context) {
		startTime := time.Now()
		counter.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
		c.JSON(200, gin.H{
			"message": "pong", //访问localhost:8080/ping页面显示{"message": "pong"}
		})
		histogram.With(prometheus.Labels{"label1": "test"}).Observe(time.Since(startTime).Seconds())
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
	// http.FileServer 实现静态文件服务
	// http.ListenAndServe(":9090", nil)
}
