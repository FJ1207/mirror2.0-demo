package handles

import (
	"demo/logic"
	"demo/tools"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

// SellWorks 售卖商品
func SellWorks(ctx *gin.Context) {
	// 接收参数
	var ReceiveCommodity *logic.ReceiveCommodity
	err := ctx.ShouldBindJSON(&ReceiveCommodity)
	logs.Info("售卖信息%v", ReceiveCommodity)
	if err != nil {
		tools.JsonFailed(ctx, 400, err.Error())
		return
	}
	// 调用函数
	err = ReceiveCommodity.SellWorks()
	if err != nil {
		tools.JsonFailed(ctx, 500, err.Error())
		return
	}
	tools.JsonSuccess(ctx, 200)
}
