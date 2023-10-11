package model

import (
	"gorm.io/gorm"
	"math/big"
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
	Collection   string `gorm:"collection"`        // 合约地址
	Creator      string `gorm:"creator"`           // 创作者
	Owner        string `gorm:"owner"`             // 所有者
	Real         int    `gorm:"real"`              // 是否存在实物，2-不存在，1-存在
	Hide         int    `gorm:"hide"`              // 是否隐藏 0-不隐藏，1-隐藏
	Mold         int    `gorm:"mold"`              // 是否铸造，0-未铸造，1-已铸造
	ContractType int    `gorm:"contract_type"`     // 合约类型 1-721,2-1155
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

// ReceiveCommodity 商品信息接收结构体
type ReceiveCommodity struct {
	TokenID      []string `json:"token_id"`   // 唯一链上ID
	Price        float64  `json:"price"`      // 价格
	Condition    int      `json:"condition"`  // 售卖状态 0-待出售，1-已出售，2-已下架
	SaleTime     string   `json:"sale_time"`  // 售卖时间
	StartTime    string   `json:"start_time"` // 开始售卖时间
	Purchaser    string   `json:"purchaser"`  // 指定购买者，为空不限制购买者
	SaleWay      int      `json:"sale_way"`   // 售卖方式 1-售卖，2-转赠，3-竞拍
	FormWorkHash string   `json:"hash"`       // 售卖hash
	Hide         int      `json:"hide"`       // 是否隐藏购买者 0-不隐藏，1-隐藏
}
type Order1 struct {
	Trader         string    // 订单创建者的地址
	Side           uint8     // 交易方向(0是买，1是卖)
	Collection     string    // NFT合约地址
	AssetType      uint8     // 资产类型(0是ERC721,1是ERC1155)
	TokenId        string    // NFT的ID
	Amount         uint8     // NFT的数量，一般是1，对于ERC1155可能有多个。
	PaymentToken   string    // 支付的代币的地址
	Price          float64   //价格
	ListingTime    time.Time //订单开始时间戳
	ExpirationTime time.Time //订单结束时间戳
	Fees           []Fee
	Salt           *big.Int
}
type SignedOrder struct {
	Order Order  `json:"order"`
	V     uint8  `json:"v"`
	R     string `json:"r"`
	S     string `json:"s"`
}

type Order struct {
	Trader         string   `json:"trader"`
	Side           uint8    `json:"side"`
	Collection     string   `json:"collection"`
	AssetType      uint8    `json:"assetType"`
	TokenId        string   `json:"tokenId"`
	Amount         string   `json:"amount"`
	Price          *big.Int `json:"price"`
	PaymentToken   string   `json:"paymentToken"`
	ListingTime    string   `json:"listingTime"`
	ExpirationTime string   `json:"expirationTime"`
	Fees           Fee      `json:"fees"`
	Salt           string   `json:"salt"`
}

type Fee struct {
	Rate      uint16 `json:"rate"`
	Recipient string `json:"recipient"`
	OrderID   uint   `json:"orderID"`
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
