package model

import (
	"gorm.io/gorm"
	"time"
)

// Work 作品表*
type Work struct {
	gorm.Model
	Name         string `grom:"name"`              // 作品名称(副本名称后加#0x)（连续创建副本）
	Content      string `gorm:"content;type:text"` // 作品存储地址
	TokenID      string `gorm:"token_id"`          // 唯一链上ID
	Works        string `gorm:"works"`             // 作品集 (第一幅的链上ID)
	Introduction string `gorm:"introduction"`      // 作品介绍
	Collection   string `gorm:"collection"`        // 所属系列地址
	Creator      string `gorm:"creator"`           // 创作者
	Owner        string `gorm:"owner"`             // 所有者
	Real         int    `gorm:"real"`              // 是否存在实物，2-不存在，1-存在
	Hide         int    `gorm:"hide"`              // 是否隐藏 0-不隐藏，1-隐藏
	Mold         int    `gorm:"mold"`              // 是否铸造，0-未铸造，1-已铸造
	ContractType int    `gorm:"contract_type"`     // 合约类型 1-721,2-1155
}

// 作品状态表*
type WorkTypes struct {
	gorm.Model
	Name    string    `gorm:"name"`     // 作品名称
	Types   string    `gorm:"types"`    // 目前状态
	TokenID string    `gorm:"token_id"` // 唯一链上ID
	Works   string    `gorm:"works"`    // 作品集 (第一幅的地址)
	Form    string    `gorm:"form"`     // 来自
	Price   float64   `gorm:"price"`    // 价格
	Time    time.Time `gorm:"time"`     // 更改时间
	Belong  int       `gorm:"belong"`   // 属于 0-创作者，1-非创作者
}

// MoldRecord  铸造记录
type MoldRecord struct {
	gorm.Model
	Name   string `gorm:"name"`   // 作品名称
	Works  string `gorm:"works"`  // 作品
	Number int    `gorm:"number"` // 铸造数量
	Hash   string `gorm:"hash"`   // hash

}

// Commodity 商品表*
type Commodities struct {
	gorm.Model
	Name         string    `grom:"name"`         // 作品名称(副本名称后加#0x)（连续创建副本）
	TokenID      string    `gorm:"token_id"`     // 唯一链上ID
	Works        string    `gorm:"works"`        // 作品集 (第一幅的地址)
	Image        string    `gorm:"image"`        // 作品Image地址
	Introduction string    `gorm:"introduction"` // 作品介绍
	Collection   string    `gorm:"collection"`   // 所属系列
	Creator      string    `gorm:"creator"`      // 创作者
	Owner        string    `gorm:"owner"`        // 所有者
	Price        float64   `gorm:"price"`        // 价格
	Conditions   int       `gorm:"conditions"`   // 售卖状态 0-待出售，1-已出售，2-已下架,3-正在购买
	StartTime    time.Time `gorm:"start_time"`   // 开始售卖时间
	SaleTime     time.Time `gorm:"sale_time"`    // 结束售卖时间
	Purchaser    string    `gorm:"purchaser"`    // 指定购买者，为空不限制购买者
	SaleWay      int       `gorm:"sale_way"`     // 售卖方式 1-购买，2-转赠，3-竞拍
	Buyers       string    `gorm:"buyers"`       // 购买者
	Hide         int       `gorm:"hide"`         // 是否隐藏购买者 0-不隐藏，1-隐藏
}

// TransactionRecord 交易记录
type TransactionRecord struct {
	gorm.Model
	Form          string `gorm:"form"`           // Form地址
	To            string `gorm:"to"`             // TO地址
	TokenId       string `gorm:"token_id"`       // 作品地址
	FormWorkHash  string `gorm:"form_work_hash"` // 出售hash
	PurchaseHash  string `gorm:"purchase_hash"`  // 购买hash
	TradingStatus int    `gorm:"trading_status"` // 交易类型,1-正在交易，2-交易结束
}

// 订单表
