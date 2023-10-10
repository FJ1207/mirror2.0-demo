/**
 * @Title 日志初始化模块
 * @Desc 该引擎使用Beego框架的log模块实现，使用方法参考以下链接
 * @Link https://www.bookstack.cn/read/beego-2.0-zh/module-logs.md
 */

package conf

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// InitLogConfig 初始化日志引擎
func InitLogConfig() {
	time := fmt.Sprintf("nodelog_%s.log", time.Now().Format("20060102"))
	var filePath string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		filePath = filepath.Join("./", "logs", time)
	} else {
		filePath = filepath.Join(".\\", "logs", time)
	}

	// 配置日志引擎
	logConfig := map[string]interface{}{
		"filename": filePath,          // 保存的文件名
		"maxlines": 100000,            // 每个文件保存的最大行数
		"maxsize":  256 * 1024 * 1024, // 每个文件保存的最大尺寸
		"daily":    true,              // 是否按照每天分割日志
		"maxdays":  15,                // 文件最多保存多少天
		"rotate":   true,              // 是否开启日志分割
		"level":    7,                 // 日志保存的时候的级别,日志保存级别：1-Alert,2-Critical,3-Error,4-Warning,5-Notice,6-Info,7-Debug;(默认值：7)
	}
	// 转换配置为字符串
	logConfStr, err := json.Marshal(logConfig)
	// 判断转换是否成功
	if err != nil {
		fmt.Println("日志引擎初始化失败(日志配置文件转换失败)：", err.Error())
		return
	}
	// 开启文件日志记录
	err = logs.SetLogger(logs.AdapterFile, string(logConfStr))
	// 判断日志引擎是否初始化成功
	if err != nil {
		logs.Error("日志引擎初始化失败：", err.Error())
		return
	}
	// 开启日志行号及文件打印
	logs.EnableFuncCallDepth(true)
	// 开启异步日志记录
	logs.Async()

	// 打印日志初始化成功
	logs.Info("日志引擎初始化成功")

	// 使用
	// logs.Debug("my book is bought in the year of ", 2016)
	// logs.Info("this %s cat is %v years old", "yellow", 3)
	// logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	// logs.Error(1024, "is a very", "good game")
	// logs.Critical("oh,crash")
}
