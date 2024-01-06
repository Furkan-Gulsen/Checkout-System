package logger

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusLogger struct {
	requestCounter *prometheus.CounterVec
	responseStatus *prometheus.CounterVec
	httpDuration   *prometheus.HistogramVec
}

func NewPrometheusLogger() *PrometheusLogger {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP Request Total.",
		},
		[]string{"path", "method"},
	)
	responseStatus := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status",
			Help: "HTTP Response Status",
		},
		[]string{"status"},
	)
	httpDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "HTTP Response Time Seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	prometheus.MustRegister(requestCounter, responseStatus, httpDuration)

	return &PrometheusLogger{
		requestCounter: requestCounter,
		responseStatus: responseStatus,
		httpDuration:   httpDuration,
	}
}

func PrometheusMiddleware(logger *PrometheusLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		start := time.Now()
		c.Next()
		duration := time.Since(start)

		status := c.Writer.Status()
		logger.requestCounter.WithLabelValues(path, method).Inc()
		logger.responseStatus.WithLabelValues(http.StatusText(status)).Inc()
		logger.httpDuration.WithLabelValues(path).Observe(duration.Seconds())
	}
}
