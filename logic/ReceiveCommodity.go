package logic

import (
	"demo/conf"
	"demo/model"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"math/big"
	"strings"
	"time"
)

var (
	URL           = "https://exchaintestrpc.okex.org"
	ChainId       = big.NewInt(65)
	ERC721address = "0x3C8C4bCEc2aA5Bbd9E2A29c014849F489370cA36"
	privateKey    = "00b4a47061a3a6c48885208bf1dd95a5fbf588e7d2bd3b35e82d1f60007dfd2c"
	countAddress  = "0x789253a85b81E246633CAd62Ae3ddBA4cABa1675"
)

// 商品上架不需要调合约
func (commodity ReceiveCommodity) SellWorks() error {

	//ethClient, err := ethclient.Dial(URL)
	//if err != nil {
	//	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	//}
	//// NFT（ERC721）合约地址：
	//contractAddress := common.HexToAddress(ERC721address)
	//// 合约和链远程客户端绑定绑定
	//contract, err := abix.NewMockERC721(contractAddress, ethClient)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 创建一个对象，用于发送交易
	//privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//TransactOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, ChainId)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//TransactOpts.GasLimit = uint64(867160 * 2)
	//TransactOpts.GasPrice = big.NewInt(65)
	//hash := new(big.Int)
	//hash.SetString(commodity.FormWorkHash, 10)
	//
	//contract.Mint(TransactOpts, common.HexToAddress(countAddress), hash)
	//// logs.Info("售卖rewTX:%v", commodity.FormWorkHash)
	//for _, item := range commodity.TokenID {
	//	value, err := PurchaseRedisGet(item)
	//	if err != nil && err.Error() != "redis: nil" {
	//		return err
	//	}
	//	if len(value) > 0 {
	//		logs.Info("售卖redis查询在售信息,%v", value)
	//		return fmt.Errorf("有艺术家正在购买您的作品，不允许重新上架此系列作品")
	//
	//	}
	//}
	// 调用上链信息
	//hash, err := contract.Mint(TransactOpts, commodity.FormWorkHash)
	//if err != nil {
	//	logs.Error("售卖调用上链进程失败，err:%v", err)
	//	return err
	//}

	// 根据TokenId 进行创建售卖商品
	for i, item := range commodity.TokenID {
		var work model.Work
		if err := conf.DB1.Model(model.Work{}).Where("token_id = ?", item).Find(&work).Error; err != nil {
			logs.Error("查询作品信息失败，err:%s", err)
			return err
		}
		var commodities int64
		if err := conf.DB1.Model(model.Commodities{}).Where("token_id = ?", item).Where("owner = ?", work.Owner).Where("conditions = ?", 2).Count(&commodities).Error; err != nil {
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
		if i == 0 {
			if err := conf.DB1.Model(model.Commodities{}).
				Where(map[string]interface{}{"works": work.Works, "owner": work.Owner, "conditions": 0}).
				Update("price", commodity.Price).Error; err != nil {
				logs.Error("查询作品信息失败，err:%s", err)
				return err
			}
		}
		// 判断是否有进行下架
		if commodities > 0 {
			if err := conf.DB1.Model(model.Commodities{}).
				Where("token_id = ?", item).
				Where("owner = ?", work.Owner).
				Where("conditions = ?", 2).
				Updates(map[string]interface{}{"sale_time": seleTime, "start_time": startTime, "price": commodity.Price, "conditions": 0}).
				Error; err != nil {
				logs.Error("查询作品信息失败，err:%s", err)
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

			// 赋值信息
			createCommodity := model.Commodities{
				Name:         work.Name,
				TokenID:      work.TokenID,
				Works:        work.Works,
				Image:        work.Content,
				Introduction: work.Introduction,
				Collection:   work.Collection,
				Creator:      work.Creator,
				Owner:        work.Owner,
				Price:        commodity.Price,
				Conditions:   commodity.Condition,
				StartTime:    startTime,
				SaleTime:     seleTime,
				Purchaser:    commodity.Purchaser,
				SaleWay:      commodity.SaleWay,
				Buyers:       "",
				Hide:         0,
			}
			// 创建作品状态信息
			result := conf.DB1.Create(&createCommodity)
			if result.Error != nil {
				logs.Info("写入商品品%s失败,时间：%v,err：%s", seleTime, createCommodity.TokenID, result.Error)
				return result.Error
			}
			logs.Info("写入商品品%s结束", createCommodity.TokenID, result.Error)
		}
		// 创建交易记录
		CreateTransactionRecord(work.Owner, commodity.FormWorkHash, work.TokenID)
		// 修改作品状态
		if strings.EqualFold(work.Creator, work.Owner) {
			UpdateWorkTypes(item, "售卖中", "我", commodity.Price, 0)
		} else {
			UpdateWorkTypes(item, "售卖中", "我", commodity.Price, 1)
		}

	}

	return nil
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
