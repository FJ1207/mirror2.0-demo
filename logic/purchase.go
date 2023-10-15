package logic

import (
	"demo/conf"
	"demo/conf/abix"
	"demo/model"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	Recipient             = "0xCd18BE3282cC6e4AB072Bab888773B6e10688888" //fee接收地址
)

// Purchase  购买
func (PurchaseCommodity *PurchaseCommodity) Purchase() error {
	//查询商品在售信息

	time := time.Now()
	var Commodities int64
	// 查找在售卖的 商品
	if err := conf.DB1.Model(model.CommoditiesDemo{}).
		Where(map[string]interface{}{"token_id": PurchaseCommodity.TokenID, "conditions": 0}).
		Where("start_time < ? AND sale_time > ?", time, time).
		Count(&Commodities).
		Error; err != nil {
		logs.Error("查找在售卖的 商品失败，err:%s", err)
		return err
	}
	if Commodities == 0 {
		return fmt.Errorf("购买失败，作品未开始售卖或已下架")
	}
	var commoditys model.CommoditiesDemo
	if err := conf.DB1.Model(model.CommoditiesDemo{}).Where(map[string]interface{}{"token_id": PurchaseCommodity.TokenID, "conditions": 0}).Where("start_time < ? AND sale_time > ?", time, time).Find(&commoditys).Error; err != nil {
		return err
	}
	tokenid, err := strconv.ParseInt(commoditys.TokenID, 10, 64)
	if err != nil {
		return err
	}
	// 方便测试，随机盐传固定值
	saltSell := "0x7a14e8fcda79940f6f4584a7c2e67433ea4245cc6e2d49134e752b6a7b6e0d23"
	saltBuy := "0x3a1f8d7e13f5f66eb840b0beabf9232a03202f9614d3c1cb6642a6f045bc1e25"

	// 支付代币地址 ：0x2219845942d28716c0F7C605765fABDcA1a7d9E0
	payaddress := "0x2219845942d28716c0F7C605765fABDcA1a7d9E0"

	saltsell, err := strconv.ParseInt(saltSell, 16, 64)
	fee := make([]abix.Fee, 0)
	//TODO :为啥是切片
	fee1 := abix.Fee{Rate: 50, Recipient: common.BytesToAddress([]byte(Recipient))}
	fee = append(fee, fee1)

	var Sellorder = abix.Order{
		Trader:         common.BytesToAddress([]byte(commoditys.Trader)),
		Side:           1,
		Collection:     common.BytesToAddress([]byte(commoditys.Collection)),
		AssetType:      0,
		TokenId:        big.NewInt(tokenid),
		Amount:         big.NewInt(1),
		Price:          big.NewInt(int64(commoditys.Price)),
		PaymentToken:   common.BytesToAddress([]byte(payaddress)),
		ListingTime:    big.NewInt(commoditys.StartTime.UnixNano()),
		ExpirationTime: big.NewInt(commoditys.SaleTime.UnixNano()),
		Fees:           fee,
		Salt:           big.NewInt(saltsell),
	}

	signature8 := "0x091e27ae0e9957c0dab85fbee814699beb15b2316e0b65f3c149f7d74e2426031da34df2970d57e0944c7805af75d9f2ad3994865093db81d1670af873ff49bc1c"

	sigHex8, err := hexutil.Decode(signature8)

	// 将查询到的商品信息形成订单
	SignedSellOrder := abix.SignedOrder{
		Order: Sellorder,
		// 签名 (实际由前端进行签名） EIP712
		V: sigHex8[64],
		R: [32]byte(sigHex8[0:32]),
		S: [32]byte(sigHex8[32:64]),
	}

	user, err := model.FindUser(PurchaseCommodity.User)
	if err != nil {
		return err
	}
	purcahseAddress := user.UserAddress

	saltbuy, err := strconv.ParseInt(saltBuy, 16, 64)

	var PurchaseOrder = abix.Order{
		Trader:         common.BytesToAddress([]byte(purcahseAddress)),
		Side:           0,
		Collection:     common.BytesToAddress([]byte(commoditys.Collection)),
		AssetType:      0,
		TokenId:        big.NewInt(tokenid),
		Amount:         big.NewInt(1),
		Price:          big.NewInt(PurchaseCommodity.Price),
		PaymentToken:   common.BytesToAddress([]byte(payaddress)),
		ListingTime:    big.NewInt(time.UnixNano()),
		ExpirationTime: big.NewInt(commoditys.SaleTime.UnixNano()),
		Fees:           fee,
		Salt:           big.NewInt(saltbuy),
	}

	signature6 := "0x397dd5b5c3c4192f799386ea0851b498d49fb5f2db7ea8d5d0bd4b27b7c28d79319d635176a3b12322041ff526eecfdb7eec99971463d96d8bcc24b373f7cf591b"

	sigHex6, err := hexutil.Decode(signature6)

	// 将查询到的商品信息形成订单
	SignedPurchaseOrder := abix.SignedOrder{
		Order: PurchaseOrder,
		// 签名 (实际由前端进行签名） EIP712
		V: sigHex6[64],
		R: [32]byte(sigHex6[0:32]),
		S: [32]byte(sigHex6[32:64]),
	}

	ethClient, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// 调用合约
	// NFT（ERC721）合约地址：
	contractAddress := common.HexToAddress(MirrorExchangeAddress)
	// 合约和链远程客户端绑定绑定
	contract, err := abix.NewMirrorExchange(contractAddress, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	// 吃单者调合约
	privateKeyECDSA, err := crypto.HexToECDSA(user.UserPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	TransactOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, ChainId)
	if err != nil {
		log.Fatal(err)
	}
	TransactOpts.GasLimit = uint64(867160 * 2)
	TransactOpts.GasPrice = big.NewInt(65)
	tx, err := contract.Execute(TransactOpts, SignedSellOrder, SignedPurchaseOrder)
	if err != nil {
		return err
	}

	// 更改作品信息
	logs.Info("交易成功hash:%v", tx.Hash().String())
	if err == nil {
		err := model.UpdateWork(PurchaseCommodity.TokenID, PurchaseCommodity.User)
		if err != nil {
			logs.Error("更改作品信息失败，err:%s", err)
			return err
		}
	}

	err1 := model.UpdatePurUser(PurchaseCommodity.User, PurchaseCommodity.Price)
	if err1 != nil {
		logs.Error("更改购买者金额记录失败，err:%s", err1)
	}
	err2 := model.UpdateSellUser(commoditys.Owner, PurchaseCommodity.Price)
	if err2 != nil {
		logs.Error("更改售卖者金额记录失败，err:%s", err2)
	}
	return nil
}

// PurchaseCommodity 购买商品
type PurchaseCommodity struct {
	TokenID string `json:"token_id"` // 唯一链上ID
	User    string `json:"user"`     // 用户名
	Price   int64  `json:"price"`    // 价格,单位(wei)
}

//type Order1 struct {
//	Trader         string    // 订单创建者的地址
//	Side           uint8     // 交易方向(0是买，1是卖)
//	Collection     string    // NFT合约地址
//	AssetType      uint8     // 资产类型(0是ERC721,1是ERC1155)
//	TokenId        string    // NFT的ID
//	Amount         uint8     // NFT的数量，一般是1，对于ERC1155可能有多个。
//	PaymentToken   string    // 支付的代币的地址
//	Price          float64   //价格
//	ListingTime    time.Time //订单开始时间戳
//	ExpirationTime time.Time //订单结束时间戳
//	Fees           []Fee1
//	Salt           *big.Int
//}
//
//type Fee1 struct {
//	Rate      uint16
//	Recipient common.Address
//}
//
//// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
//type SignedOrder1 struct {
//	Order Order1
//	V     uint8
//	R     string
//	S     string
//}

func GenerateSalt() string {
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
