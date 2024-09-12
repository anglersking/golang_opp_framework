package main

import (
	"fmt"
	"DDD/sample/sample_BLL"
	"github.com/sirupsen/logrus"
)

func main() {
	Log := logrus.New()

	// 设置日志格式为 JSON
	// log.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志格式为文本格式
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // 显示完整的时间戳
	})

	// 设置日志级别为 Info
	Log.SetLevel(logrus.InfoLevel)
	Log.Info("This is an info message")
	Log.Warn("This is a warning message")
	Log.Error("This is an error message")

	// // 带有字段的日志
	// log.WithFields(logrus.Fields{
	// 	"event": "event_name",
	// 	"topic": "topic_name",
	// }).Info("A structured log message")

	testManager := bll.NewTestManager()

	// 获取模型值
	Log.Info("Initial Model Value:", testManager.GetModelValue())

	// 设置新的模型值
	testManager.SetModelValue(456)
	fmt.Println("Updated Model Value:", testManager.GetModelValue())
}
