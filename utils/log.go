package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log *Logger

type Logger struct {
	// 日志目录
	LogPath string `yaml:"logpath"`
	// 日志文件名称前缀
	LogPrefix string `yaml:"logprefix"`
	// 是否开启 请求的 trace
	IsTrace bool `yaml:"trace"`
	// 日志级别
	Level uint32 `yaml:"level"`
	// 日志格式 text json
	Formatter string `yaml:"formatter"`
	Logrus    *logrus.Logger
}

// 初始化日志
func InitLogger(l *Logger) {
	if err := os.MkdirAll(l.LogPath, 0777); err != nil {
		fmt.Println("日志目录创建失败：" + err.Error())
		os.Exit(1)
	}
	if len(l.Formatter) < 1 {
		l.Formatter = "text"
	}
	if l.Formatter != "text" && l.Formatter != "json" {
		fmt.Println("日志滚动参数配置错误")
		os.Exit(1)
	}

	logFileName := l.LogPrefix + ".log"

	fileName := path.Join(l.LogPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println("日志文件创建失败：" + err.Error())
			os.Exit(1)
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("日志文件打开失败：" + err.Error())
		os.Exit(1)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.Level(l.Level))

	switch l.Formatter {
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	l.Logrus = logger
	// l.Logrus.SetReportCaller(true)
	l.Logrus.AddHook(l)
	Log = l
}

func (l *Logger) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *Logger) Fire(entry *logrus.Entry) error {
	entry.Data["hostname"] = HostName()
	return nil
}

func (l *Logger) Ix(c *gin.Context, tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"trace_id": c.GetString("traceid"),
		"tag":      tag,
	}).Info(msg)
}

func (l *Logger) I(tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"tag": tag,
	}).Info(msg)
}

func (l *Logger) Wx(c *gin.Context, tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"trace_id": c.GetString("traceid"),
		"tag":      tag,
	}).Warn(msg)
}

func (l *Logger) W(tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"tag": tag,
	}).Warn(msg)
}

func (l *Logger) Fx(c *gin.Context, tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"trace_id": c.GetString("traceid"),
		"tag":      tag,
	}).Fatal(msg)
}

func (l *Logger) F(tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"tag": tag,
	}).Fatal(msg)
}

func (l *Logger) Ex(c *gin.Context, tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"trace_id": c.GetString("traceid"),
		"tag":      tag,
	}).Error(msg)
}

func (l *Logger) E(tag string, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	l.Logrus.WithFields(logrus.Fields{
		"tag": tag,
	}).Error(msg)
}
