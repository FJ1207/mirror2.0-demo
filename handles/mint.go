package handles

import (
	"demo/logic"
	"demo/tools"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

// CreateMold 铸造
// @param rew_tx        string    rewTX
// @param number        string    铸造数量
// @param works         string    rewTX
// @return   /super/nft/front/query/ortfoliSale [GET]
func CreateMold(ctx *gin.Context) {
	var MintReq *logic.MintReq
	err := ctx.ShouldBindJSON(&MintReq)
	logs.Info("铸造 : %v", MintReq)
	if err != nil {
		tools.JsonFailed(ctx, 400, err.Error())
		return
	}
	err = MintReq.Mint()
	if err != nil {
		tools.JsonFailed(ctx, 500, err.Error())
		return
	}
	tools.JsonSuccess(ctx, 200)
}
