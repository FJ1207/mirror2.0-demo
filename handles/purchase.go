package handles

import (
	"demo/logic"
	"demo/tools"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Purchase 购买商品
// @param  form_address    string   form地址信息
// @param  to_address      string   to地址信息
// @param  work_address    string   作品地址信息
// @param  purchase_hash   string   购买hash
// @param  pledge          string   金额
// @return /super/nft/worm/purchase [POST]
func Purchase(ctx *gin.Context) {
	formAddress := ctx.Query("from_address")
	toAddress := ctx.Query("to_address")
	tokenId := ctx.Query("work_address")
	purchaseHash := ctx.Query("hash")
	pledgeStr := ctx.Query("pledge")
	logs.Info("调用购买商品")
	if len(tokenId) == 0 || len(formAddress) == 0 || len(toAddress) == 0 || len(pledgeStr) == 0 || len(purchaseHash) == 0 {
		tools.JsonFailed(ctx, 100, "缺少必要参数")
		return
	}
	pledge, err := strconv.ParseFloat(pledgeStr, 64)
	if err != nil {
		tools.JsonFailed(ctx, 100, "类型转换失败")
		return
	}

	err = logic.Purchase(formAddress, toAddress, tokenId, purchaseHash, pledge)
	if err != nil {
		tools.JsonFailed(ctx, 500, err.Error())
		return
	}
	tools.JsonSuccess(ctx, 200)
}
