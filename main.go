package main

import (
	"demo/conf"
	"demo/handles"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"os"
)

// Init 初始化进程
func Init() {
	// 初始化日志
	conf.InitLogConfig()
	// 初始化数据库
	conf.InitMySQLDatabase()

	// 初始化 客户端
	//conf.InitChainClient()

	// 创建日志文件夹
	err := os.MkdirAll("./logs", 0755)
	if err != nil {
		s := fmt.Errorf("create dir \"%s\" error: %s", "./logs", err.Error())
		panic(s)
	}

}

func main() {

	// 创建一个无中间件的路由
	r := gin.New()
	// 判断是否启用日志
	if gin.Mode() == gin.DebugMode {
		// 启用日志打印
		r.Use(gin.Logger())
	}
	// 初始化路由
	handles.InitRouter(r)
	// 拼接监听地址
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 9999)
	// 记录启动日志
	logs.Info("HTTP服务将在以下地址开启监听：\"http://%s\"", address)
	// 启动服务
	if err := r.Run(address); err != nil {
		logs.Error("HTTP服务启动失败")
	}
}
