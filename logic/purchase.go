package logic

import (
	"demo/conf"
	"demo/conf/abix"
	"demo/model"
	"demo/tools"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

var (
	MirrorExchangeAddress = "0x50047A8658baBADA196C395589A8Fc03178fD282"
	Recipient ="0xCd18BE3282cC6e4AB072Bab888773B6e10688888"
)
// Purchase  购买
func (PurchaseCommodity PurchaseCommodity)Purchase() error {
	//查询商品在售信息

	time := time.Now()
	var Commodities int64
	// 查找在售卖的 商品
	if err := conf.DB1.Model(model.CommoditiesDemo{}).
		Where(map[string]interface{}{"token_id": PurchaseCommodity.TokenID, "conditions": 0}).
		Where("start_time < ? AND sale_time > ?", time, time).
		Count(&Commodities).
		Error; err != nil {
		logs.Error("转赠商品失败，err:%s", err)
		return err
	}
	if Commodities == 0 {
		return fmt.Errorf("购买失败，作品未开始售卖或已下架")
	}
	var commoditys model.CommoditiesDemo
	if err := conf.DB1.Model(model.CommoditiesDemo{}).Where(map[string]interface{}{"token_id": PurchaseCommodity.TokenID, "conditions": 0}).Where("start_time < ? AND sale_time > ?", time, time).Find(&commoditys).Error; err != nil {
		return err
	}
	tokenid ,err := strconv.ParseInt(commoditys.TokenID,10,64)
	if err != nil {
		return err
	}
	salt, err := strconv.ParseInt(GenerateSalt(), 16, 64)
	fee := make([]Fee,0)
	//TODO :为啥是切片
	fee1 := Fee{Rate: 50,Recipient:common.BytesToAddress([]byte(Recipient)) }
	fee = append(fee,fee1)

	var order = Order{
		Trader: common.BytesToAddress([]byte(commoditys.Trader)),
		Side: 1,
		Collection: common.BytesToAddress([]byte(commoditys.Collection)),
		AssetType: 0,
		TokenId : big.NewInt(tokenid),
		Amount: big.NewInt(1),
		//TODO:合约里面的价格是整数？
		Price: big.NewInt(int64(commoditys.Price)),
		PaymentToken: common.BytesToAddress([]byte{}),
		ListingTime: big.NewInt(commoditys.StartTime.UnixNano()),
		ExpirationTime: big.NewInt(commoditys.SaleTime.UnixNano()),
		Fees : fee,
		Salt : big.NewInt(salt),
	}
	v, r, s, err := tools.Sign(order)
	if err != nil {
		return err
	}

	var arrR [32]byte
	copy(arrR[:], r)

	var arrS [32]byte
	copy(arrS[:], s)

	// 将查询到的商品信息形成订单
	SignedSellOrder :=  SignedOrder{
		Order: order,
		// 签名 (实际由前端进行签名） EIP712
		V: v,
		R: arrR,
		S: arrS,
	}

	// 形成买单::wq:
	var SignedPurchaseOrder conf.
	SignedPurchaseOrder =  SignedOrder{
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
	privateKeyECDSA, err := crypto.HexToECDSA(PrivateKey)
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

// PurchaseCommodity 购买商品
type PurchaseCommodity struct {
	toAddress string  `json:"to_address"` // to地址信息
	TokenID      string  `json:"token_id"`   // 唯一链上ID
	hash   string  `json:"hash"` // 购买hash
	Price        float64 `json:"price"`      // 价格
}


type Order1 struct {
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
	Fees           []Fee1
	Salt           *big.Int
}

type Fee1 struct {
	Rate      uint16
	Recipient common.Address
}

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder1 struct {
	Order Order
	V     uint8
	R     string
	S     string
}

type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	Collection     common.Address
	AssetType      uint8
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
}

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder struct {
	Order Order
	V     uint8
	R     [32]byte
	S     [32]byte
}

func GenerateSalt() (string) {
	var array [8]uint32
	for i := 0; i < 8; i++ {
		array[i] = rand.Uint32()
	}
	hexString := ""
	for i := 0; i < len(array); i++ {
		hexString += fmt.Sprintf("%08x", array[i])
	}

	return hexString
}