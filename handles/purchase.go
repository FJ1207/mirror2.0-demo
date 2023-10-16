package handles

import (
	"demo/logic"
	"demo/tools"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"runtime"
)

// Purchase 购买商品
// @param  form_address    string   form地址信息
// @param  to_address      string   to地址信息
// @param  work_address    string   作品地址信息
// @param  purchase_hash   string   购买hash
// @param  pledge          string   金额
// @return /super/nft/worm/purchase [POST]
func Purchase(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// Print the stack trace
			var stacktrace string
			for i := 0; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", file, line)
				fn := runtime.FuncForPC(pc)
				if fn != nil {
					stacktrace += fmt.Sprintf("  %s\n", fn.Name())
				}
			}
			logs.Info(fmt.Sprintf("panic error: %v\n%s", r, stacktrace))
		}
	}()
	// 形成买单
	var PurchaseCommodity *logic.PurchaseCommodity
	err := ctx.ShouldBindJSON(&PurchaseCommodity)
	logs.Info("购买商品 : %v", PurchaseCommodity)
	if err != nil {
		tools.JsonFailed(ctx, 400, err.Error())
		return
	}
	err = PurchaseCommodity.Purchase()
	if err != nil {
		tools.JsonFailed(ctx, 500, err.Error())
		return
	}
	tools.JsonSuccess(ctx, 200)
}

// 生成随机盐
