package main

import (
	"go.uber.org/zap"
)

func main()  {
	logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	url := "http://127.0.0.1"
	//方式一：性能极致
	logger.Info("fail to fetch url",
		zap.String("url",url),
		zap.Int("nums",3),
		)
	//方式二：性能比一般的日志库快4-10倍
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	//"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}
