/**
 * @Title 返回前端的数据格式
 * @Desc 再次约束返回给前端的数据格式
 */

package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 数据返回结构
type ResponseData struct {
	Code uint32      `json:"code"`           // 状态码
	Msg  string      `json:"msg"`            // 状态描述
	Data interface{} `json:"data,omitempty"` // 常规json数据或PaginationData数据
}

// PaginationData 分页数据格式
type PaginationData struct {
	TotalData  int64       `json:"total_data"`  // 数据总量
	TotalPages int         `json:"total_pages"` // 总页数
	CurrPage   int         `json:"curr_page"`   // 当前页码
	MaxNum     int         `json:"max_num"`     // 约束的最大数据量
	Data       interface{} `json:"data"`        // 数据
}

// JsonSuccess 成功结果返回
// @params ctx  *gin.Context Gin框架上下文
// @params data interface{}  返回数据
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.Status(http.StatusOK)
	ctx.JSON(200, ResponseData{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
	ctx.Abort()
}

// JsonFailed 失败结果返回
// @params ctx     *gin.Context Gin框架上下文
// @params code    uint32       自定义状态码
// @params msg     string       状态描述
func JsonFailed(ctx *gin.Context, code uint32, msg string) {
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  msg,
	})
	ctx.Abort()
}
