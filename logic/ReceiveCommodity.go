package logic

import (
	"demo/conf"
	"demo/model"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"strings"
	"time"
)

// 商品上架不需要调合约
func (commodity *ReceiveCommodity) SellWorks() error {
	var work model.WorkDemo
	if err := conf.DB1.Model(model.WorkDemo{}).Where("token_id = ?", commodity.TokenID).Find(&work).Error; err != nil {
		logs.Error("查询作品信息失败，err:%s", err)
		return err
	}
	var commodities int64
	if err := conf.DB1.Model(model.CommoditiesDemo{}).Where("token_id = ?", commodity.TokenID).Where("owner = ?", work.Owner).Where("conditions = ?", 2).Count(&commodities).Error; err != nil {
		logs.Error("查询作品信息失败，err:%s", err)
		return err
	}
	sele := strings.ReplaceAll(commodity.SaleTime, "/", "-") // 替换全部
	seleTime, err := time.ParseInLocation(`2006-1-2 15:04:05`, sele, time.Local)
	if err != nil {
		logs.Error("转换时间失败,err：%v,,,%s", commodity.SaleTime, err)
		return fmt.Errorf("上架时间错误，请联系相关人员询问，或稍后再试")
	}
	start := strings.ReplaceAll(commodity.StartTime, "/", "-") // 替换全部
	startTime, err := time.ParseInLocation(`2006-1-2 15:04:05`, start, time.Local)
	if err != nil {
		logs.Error("转换时间失败,err：%v,%s", commodity.StartTime, err)
		return fmt.Errorf("上架时间错误，请联系相关人员询问，或稍后再试")
	}

	//if err := conf.DB1.Model(model.Commodities{}).
	//	Where(map[string]interface{}{"owner": work.Owner, "conditions": 0}).
	//	Update("price", commodity.Price).Error; err != nil {
	//	logs.Error("查询作品信息失败，err:%s", err)
	//	return err
	//}

	// 判断是否有进行下架
	if commodities > 0 {
		if err := conf.DB1.Model(model.Commodities{}).
			Where("token_id = ?", commodity.TokenID).
			Where("owner = ?", work.Owner).
			Where("conditions = ?", 2).
			Updates(map[string]interface{}{"sale_time": seleTime, "start_time": startTime, "price": commodity.Price, "conditions": 0}).
			Error; err != nil {
			logs.Error("商品重新上架失败，err:%s", err)
			return err
		}
	} else {
		// 判断是否是实体作品
		if work.Real == 1 {
			bools := strings.EqualFold(work.Creator, work.Owner)
			if !bools {
				return fmt.Errorf("实体作品只能售卖一次")
			}
		}
		user, err := model.FindUser(work.Owner)
		if err != nil {
			return err
		}
		// 没有下架，也不是售卖一次的实体作品，创建商品信息
		createCommodity := model.CommoditiesDemo{
			Trader:     user.UserAddress,
			TokenID:    work.TokenID,
			Collection: work.Collection,
			Price:      commodity.Price,
			Conditions: 0,
			StartTime:  startTime,
			SaleTime:   seleTime,
			Creator:    work.Creator,
			Owner:      work.Owner,
		}
		// 创建作品状态信息
		result := conf.DB1.Create(&createCommodity)
		if result.Error != nil {
			logs.Info("写入商品品%s失败,时间：%v,err：%s", seleTime, createCommodity.TokenID, result.Error)
			return result.Error
		}
		logs.Info("写入商品品%s结束", createCommodity.TokenID, result.Error)
	}
	//// 创建交易记录
	//CreateTransactionRecord(work.Owner, commodity.FormWorkHash, work.TokenID)
	//// 修改作品状态
	//if strings.EqualFold(work.Creator, work.Owner) {
	//	UpdateWorkTypes(commodity.TokenID, "售卖中", "我", commodity.Price, 0)
	//} else {
	//	UpdateWorkTypes(commodity.TokenID, "售卖中", "我", commodity.Price, 1)
	//}
	return nil
}

// ReceiveCommodity 商品信息接收结构体
type ReceiveCommodity struct {
	//Trader       string  `json:"trader"`     //创建订单者地址
	TokenID string `json:"token_id"` // 唯一链上ID
	//Collection   string  `json:"collection"` // 作品合约地址
	Price int64 `json:"price"` // 价格，单位wei
	//Condition    int     `json:"condition"`  // 售卖状态 0-待出售，1-已出售，2-已下架
	SaleTime  string `json:"sale_time"`  // 结束售卖时间
	StartTime string `json:"start_time"` // 开始售卖时间
	//FormWorkHash string  `json:"hash"`       // 售卖hash
}
