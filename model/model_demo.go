package model

import (
	"gorm.io/gorm"
	"time"
)

// 用户表
type UserDemo struct {
	gorm.Model
	Name           string `grom:"name"`             // 用户名称
	UserAddress    string `gorm:"user_address"`     // 用户地址
	UserPrivateKey string `grom:"user_private_key"` //用户私钥
	Amount         int64  `grom:"amount"`           // 金额
}

// nft表
type WorkDemo struct {
	gorm.Model
	Name         string `grom:"name"`                // 作品名称(副本名称后加#0x)（连续创建副本）
	Content      string `gorm:"content;type:text"`   // 作品存储地址
	TokenID      string `gorm:"token_id;primaryKey"` // 唯一链上ID
	Introduction string `gorm:"introduction"`        // 作品介绍
	Creator      string `gorm:"creator"`             // 创作者
	Owner        string `gorm:"owner"`               // 所有者
	Real         int    `gorm:"real"`                // 是否存在实物，2-不存在，1-存在
	AssetType    int    `gorm:"asset_type"`          // 资产类型(0-ERC721,1-ERC1155)
	Collection   string `gorm:"collection"`          // 作品合约地址
}

// 商品表--可以转化为卖单
type CommoditiesDemo struct {
	gorm.Model
	//Side       int       `gorm:"side"`                // 交易方向(0是买，1是卖)

	Trader     string    `grom:"trader"`              // 订单创建者的地址
	TokenID    string    `gorm:"token_id;primaryKey"` // 唯一链上ID
	Collection string    `gorm:"collection"`          // 作品合约地址
	Price      int64     `gorm:"price"`               // 价格
	Conditions int       `gorm:"conditions"`          // 售卖状态 0-待出售，1-已出售，2-已下架,3-正在购买
	StartTime  time.Time `gorm:"start_time"`          // 开始售卖时间
	SaleTime   time.Time `gorm:"sale_time"`           // 结束售卖时间
	//Salt       string    `gorm:"salt"`                //随机盐
	Creator string `gorm:"creator"` // 创作者
	Owner   string `gorm:"owner"`   // 所有者
}

// 作品状态表*
type WorkTypesDemo struct {
	gorm.Model
	Name    string    `gorm:"name"`     // 作品名称
	Types   string    `gorm:"types"`    // 目前状态
	TokenID string    `gorm:"token_id"` // 唯一链上ID
	Form    string    `gorm:"form"`     // 来自
	Price   float64   `gorm:"price"`    // 价格
	Time    time.Time `gorm:"time"`     // 更改时间
	Belong  int       `gorm:"belong"`   // 属于 0-创作者，1-非创作者
}
