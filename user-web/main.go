package main

import (
	"fmt"
	"xshop-api/user-web/initialize"

	"go.uber.org/zap"
)

func main() {
	port := 8021
	//初始化日志
	initialize.InitLogger()
	//初始化路由
	Router := initialize.Routers()

	/*
		1. S() 可以获取 一个全局的sugar,可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug,info,warn,error,fetal
		3.S函数和L函数很有用，提供一个全局的安全访问logger的访问
	*/
	zap.S().Infof("启动服务器，端口：%d", port)

	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

}
