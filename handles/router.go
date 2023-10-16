package handles

import (
	"demo/tools"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	// 为所有路由添加Recover异常捕获（返回自己定义的内容）
	router.Use(AppRecover)
	router.HEAD("/head", func(ctx *gin.Context) {
		tools.JsonSuccess(ctx, nil)
	})
	// 定义api路由分组
	nft := router.Group("/nft")
	work := nft.Group("/worm")
	{
		work.POST("/mold", CreateMold)      // 铸造
		work.POST("/sell/works", SellWorks) // 售卖商品
		work.POST("/purchase", Purchase)    // 购买商品
	}

}

// midware
func AppRecover(ctx *gin.Context) {
	// 全局捕获异常
	defer func() {
		if err := recover(); err != nil {
			logs.Error("服务器发生异常错误：%s", fmt.Sprint(err))
			tools.JsonFailed(ctx, 500, "网络异常")
		}
	}()
	// 调用下级
	ctx.Next()
}
