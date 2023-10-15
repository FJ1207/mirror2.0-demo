package logic

import (
	"context"
	"demo/conf/abix"
	"demo/model"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	URL           = "https://exchaintestrpc.okex.org"
	ChainId       = big.NewInt(65)
	ERC721address = "0x3C8C4bCEc2aA5Bbd9E2A29c014849F489370cA36"
	PrivateKey    = "00b4a47061a3a6c48885208bf1dd95a5fbf588e7d2bd3b35e82d1f60007dfd2c"
	countAddress  = "0x789253a85b81E246633CAd62Ae3ddBA4cABa1675"
)

func (MintReq *MintReq) Mint() error {

	// 使用 SetString 将字符串转换为 *big.Int
	tokenID := new(big.Int)
	tokenID.SetString(MintReq.TokenId, 10)

	ethClient, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return err
	}
	// NFT（ERC721）合约地址：
	contractAddress := common.HexToAddress(ERC721address)
	// 合约和链远程客户端绑定绑定
	contract, err := abix.NewMockERC721(contractAddress, ethClient)
	if err != nil {
		log.Fatal(err)
		return err
	}
	user, err := model.FindUser(MintReq.User)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 创建一个对象，用于发送交易
	privateKeyECDSA, err := crypto.HexToECDSA(user.UserPrivateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	TransactOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, ChainId)
	if err != nil {
		log.Fatal(err)
		return err
	}
	TransactOpts.GasLimit = uint64(867160 * 2)
	TransactOpts.GasPrice = big.NewInt(65)

	tx, err := contract.Mint(TransactOpts, common.HexToAddress(countAddress), tokenID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	hash := tx.Hash().String()
	fmt.Printf("tx : %v ", tx)

	// 等待交易被确认
	receipt, err := bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 检查交易是否成功
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("Minting NFT failed")
	} else {
		log.Println("NFT minted successfully")
	}
	var workDemo = model.WorkDemo{
		Name:         MintReq.TokenId,
		TokenID:      MintReq.TokenId,
		Content:      hash,
		Introduction: MintReq.Tntroduction,
		Real:         MintReq.Real,
		AssetType:    MintReq.AssetType,
		Creator:      MintReq.User,
		Owner:        MintReq.User,
		Collection:   ERC721address,
	}

	if err == nil {
		if err := model.CreateWork(&workDemo); err != nil {
			logs.Error("创建铸造信息错误，err:%s", err)
			return err
		}
	}
	return nil
}

type MintReq struct {
	TokenId      string `json:"token_id"`
	Tntroduction string `json:"introduction"`
	Real         int    `json:"real"`
	AssetType    int    `json:"asset_type"`
	User         string `json:"user"` //用户名
}
