package logic

import (
	"demo/conf"
	"demo/model"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

func CreateTransactionRecord(form, formWorkHash, tokenId string) {
	// 赋值
	transactionRecord := &model.TransactionRecord{
		Form:          form,
		TokenId:       tokenId,
		FormWorkHash:  formWorkHash,
		TradingStatus: 1,
	}
	result := conf.DB1.Create(&transactionRecord)
	if result.Error != nil {
		logs.Error("创建交易记录结束,：%s", result.Error)
		return
	}
}

// UpdateWorkTypes  更新作品状态信息
// @param works   string   作品id
func UpdateWorkTypes(work, types, form string, price float64, belong int) error {
	logs.Info("更新作品状态信息")
	time := time.Now()
	// 更新作品状态信息
	if err := conf.DB1.Model(model.WorkTypesDemo{}).
		Where("token_id = ?", work).
		Where("belong = ?", belong).
		Updates(map[string]interface{}{"types": types, "price": price, "time": time, "form": form}).
		Error; err != nil {
		logs.Error("更新作品状态信息失败，err:%s", err)
		return err
	}
	return nil
}
