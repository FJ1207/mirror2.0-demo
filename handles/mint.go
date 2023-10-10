package handles

import (
	"context"
	"demo/abix"
	"demo/conf"
	"demo/model"
	"demo/tools"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strconv"
)

var (
	URL           = "https://exchaintestrpc.okex.org"
	ChainId       = big.NewInt(65)
	ERC721address = "0x3C8C4bCEc2aA5Bbd9E2A29c014849F489370cA36"
	privateKey    = "00b4a47061a3a6c48885208bf1dd95a5fbf588e7d2bd3b35e82d1f60007dfd2c"
	countAddress  = "0x789253a85b81E246633CAd62Ae3ddBA4cABa1675"
)

// CreateMold 铸造
// @param rew_tx        string    rewTX
// @param number        string    铸造数量
// @param works         string    rewTX
// @return   /super/nft/front/query/ortfoliSale [GET]
func CreateMold(ctx *gin.Context) {
	works := ctx.Query("works")
	number, err := strconv.Atoi(ctx.Query("number"))
	tokenIDStr := ctx.Query("token_id")

	// 使用 SetString 将字符串转换为 *big.Int
	tokenID := new(big.Int)
	tokenID.SetString(tokenIDStr, 10)

	ethClient, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// NFT（ERC721）合约地址：
	contractAddress := common.HexToAddress(ERC721address)
	// 合约和链远程客户端绑定绑定
	contract, err := abix.NewMockERC721(contractAddress, ethClient)
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

	tx, err := contract.Mint(TransactOpts, common.HexToAddress(countAddress), tokenID)
	if err != nil {
		log.Fatal(err)
	}
	hash := tx.Hash().String()
	fmt.Printf("tx : %v ", tx)

	// 等待交易被确认
	receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatal(err)
	}
	// 检查交易是否成功
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("Minting NFT failed")
	} else {
		log.Println("NFT minted successfully")
	}
	if err == nil {
		if err := conf.DB1.Model(model.Work{}).Where("works = ?", works).Update("mold", 1).Error; err != nil {
			logs.Error("更新铸造信息错误，err:%s", err)
		}
		// 更新作品状态
		model.UpdateWorksTypes(works, "已铸造", "我", 0, 0)
		// 创建铸造记录
		err1 := model.CreateMoldRecord(works, hash, number)
		if err1 != nil {
			log.Fatal(err)
		}
	}

	tools.JsonSuccess(ctx, 200)
}
