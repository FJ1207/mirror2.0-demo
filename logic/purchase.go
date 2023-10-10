package logic

import (
	"demo/abix"
	"demo/conf"
	"demo/model"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"log"
	"math/big"
	"strings"
	"time"
)

var (
	MirrorExchangeAddress = "0x50047A8658baBADA196C395589A8Fc03178fD282"
)
// Purchase  购买
func Purchase(formAddress, toAddress, tokenId, purchaseHash string, pledgeStr float64) error {
	//查询商品在售信息

	time := time.Now()
	var Commodities int64
	// 查找在售卖的 商品
	if err := conf.DB1.Model(model.Commodities{}).
		Where(map[string]interface{}{"token_id": tokenId, "conditions": 0}).
		Where("start_time < ? AND sale_time > ?", time, time).
		Count(&Commodities).
		Error; err != nil {
		logs.Error("转赠商品失败，err:%s", err)
		return err
	}
	if Commodities == 0 {
		return fmt.Errorf("购买失败，作品未开始售卖或已下架")
	}
	var commodity model.Commodities
	if err := conf.DB1.Model(model.Commodities{}).Where(map[string]interface{}{"token_id": tokenId, "conditions": 0}).Where("start_time < ? AND sale_time > ?", time, time).Find(&commodity).Error; err != nil {
		return err
	}

	// 将查询到的商品信息形成订单
	SignedSellOrder :=  SignedOrder{
		Order: Order{
			Trader: formAddress,
			Side: 0,
			Collection: ERC721address,
			AssetType: 0,
			TokenId : tokenId,
			Amount: 1,
			Price: big.NewInt(commodity.Price.(int64)),
			PaymentToken: "",
			ListingTime: commodity.StartTime,
			ExpirationTime: commodity.SaleTime,
			Fees : Fee{},
			salt : 0
		},
		// 签名 (实际由前端进行签名） EIP712
		V: v,
		R: r,
		S: s,
	}

	// 形成买单
	SignedPurchaseOrder :=  SignedOrder{
		Order: Order{
			Trader: formAddress,
			Side: 1,
			Collection: ERC721address,
			AssetType: 0,
			TokenId : tokenId,
			Amount: 1,
			Price: pledgeStr,
			PaymentToken: "",
			ListingTime: time,
			ExpirationTime: time,
			Fees : Fee{},
			salt : 0
		},
		// 签名 (实际由前端进行签名） EIP712
		V: v,
		R: r,
		S: s,
	}

	ethClient, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// 调用合约
	// NFT（ERC721）合约地址：
	contractAddress := common.HexToAddress(MirrorExchangeAddress)
	// 合约和链远程客户端绑定绑定
	contract, err := abix.NewMirrorExchange(contractAddress,ethClient)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个对象，用于发送交易
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	TransactOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, ChainId)
	if err != nil {
		log.Fatal(err)
	}
	TransactOpts.GasLimit = uint64(867160 * 2)
	TransactOpts.GasPrice = big.NewInt(65)
	tx ,err := contract.Execute(TransactOpts,SignedSellOrder,SignedPurchaseOrder)
	if err != nil{
		return err
	}

	// 更改作品信息
	logs.Info("购买更改作品信息rewTX:%v", purchaseHash)
	if err == nil {
		err := UpdateWork(formAddress, toAddress, tokenId, hash, pledgeStr)
		if err != nil {
			logs.Error("更改作品信息失败，err:%s", err)
			return err
		}
	}
	err1 := WriteBuyerRecord(toAddress, pledgeStr)
	if err1 != nil {
		logs.Error("更改购买者金额记录失败，err:%s", err1)
	}
	err2 := WriteSellerRecord(formAddress, pledgeStr)
	if err2 != nil {
		logs.Error("更改售卖者金额记录失败，err:%s", err2)
	}
	return nil
}


type Order struct {
	Trader         string // 订单创建者的地址
	Side           uint8          // 交易方向(0是买，1是卖)
	Collection     string // NFT合约地址
	AssetType      uint8          // 资产类型(0是ERC721,1是ERC1155)
	TokenId        string       // NFT的ID
	Amount          uint8     // NFT的数量，一般是1，对于ERC1155可能有多个。
	PaymentToken   string // 支付的代币的地址
	Price          float64       //价格
	ListingTime    time.Time     //订单开始时间戳
	ExpirationTime  time.Time       //订单结束时间戳
	Fees           []Fee
	Salt           *big.Int
}

type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder struct {
	Order Order
	V     uint8
	R     string
	S     string
}