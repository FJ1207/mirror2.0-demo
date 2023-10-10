package conf

//
//import (
//	"demo/abix"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"log"
//	"math/big"
//)
//
//var (
//	URL           = "https://exchaintestrpc.okex.org"
//	ChainId       = big.NewInt(65)
//)
//
//func InitChainClient(address string) {
//	ethClient, err := ethclient.Dial(URL)
//	if err != nil {
//		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
//	}
//	// NFT（ERC721）合约地址：
//	contractAddress := common.HexToAddress(address)
//	// 合约和链远程客户端绑定绑定
//	contract, err := abix.NewMockERC721(contractAddress, ethClient)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
