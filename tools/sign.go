package tools

import (
	"demo/conf/abix"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// 订单签名
func Sign(order abix.Order) (uint8, string, string, error) {
	//// 设置 EIP-712 域信息
	//domain := apitypes.TypedDataDomain{
	//	Name:              "Mirror Exchange",                            // 合约名称，填"Mirror Exchange"
	//	Version:           "1.0",                                        // 版本，填"1.0"
	//	ChainId:           math.NewHexOrDecimal256(65),                  // 合约所在链的chainID
	//	VerifyingContract: "0x50047A8658baBADA196C395589A8Fc03178fD282", // 已经部署好的合约地址
	//}
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		return 0, "", "", err
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(order)
	if err != nil {
		// 处理错误
		panic(err)
	}
	// 使用私钥进行签名
	signature, err := crypto.Sign(jsonData, privateKey)
	if err != nil {
		// 处理错误
		panic(err)
	}
	// 分离 v, r, 和 s
	sigV := uint8(signature[64])     // 将 v 转为 uint8
	sigR := string(signature[:32])   // 截取 r
	sigS := string(signature[32:64]) // 截取 s

	return sigV, sigR, sigS, nil
}
