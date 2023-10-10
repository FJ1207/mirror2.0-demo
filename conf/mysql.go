/**
 * @Title MySQL数据库初始化模块
 * @Desc MySQL初始化连接将在这里完成
 */

package conf

import (
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MyDB 全局MySQL数据库操作对象
var (
	DB1 *gorm.DB // Mirror2-demo-test 数据库
)

// initMySQLDatabase 初始化MySQL数据库
func InitMySQLDatabase() {
	db1, err := gorm.Open(mysql.Open("root:blx@2022@tcp(154.85.62.206:63306)/mirror2_demo_test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database2")
	}
	DB1 = db1
	logs.Info("MySQL数据库初始化连接成功")
}
