package routers

import (
	"bytes"
	"fmt"
	"gintest/utils"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// 设置traceid
func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceid := c.GetHeader("traceid")
		if traceid == "" {
			traceid = uuid.NewV4().String()
		}
		c.Set("traceid", traceid)
		c.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		utils.Log.Logrus.WithFields(logrus.Fields{
			"uri":             c.Request.RequestURI,
			"trace_id":        c.GetString("traceid"),
			"latency":         fmt.Sprintf("%.3f", float64(time.Since(startTime).Microseconds())/1000),
			"status":          c.Writer.Status(),
			"client_ip":       c.ClientIP(),
			"method":          c.Request.Method,
			"user_agent":      c.Request.UserAgent(),
			"response_length": c.Writer.Size(),
			"param":           c.Request.URL.RawQuery,
			"response":        blw.body.String(),
			"tag":             "trace",
		}).Info()
	}

}
