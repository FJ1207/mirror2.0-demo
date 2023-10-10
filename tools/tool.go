/**
 * @Title 工具模块
 * @Desc 全部小工具将在这里定义完成
 */

package tools

import (
	"archive/zip"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

// GinGetImageByte 获取Gin框架接收到的图片流
// @params ctx      *gin.Context Gin框架上下文指针
// @params label    string       表单字段名
// @return          []byte       图片流
// @return          string       文件名
// @return          error        错误信息
func GinGetImageByte(ctx *gin.Context, label string) ([]byte, string, error) {
	// 接收图像资源
	fileHeader, err := ctx.FormFile(label)
	if err != nil {
		return nil, "", err
	}
	// 获取文件名后缀
	fileSuffix := path.Ext(fileHeader.Filename)
	// 受支持的文件后缀
	var suffixName = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	// 判断文件类型
	if _, ok := suffixName[fileSuffix]; !ok {
		return nil, "", fmt.Errorf("文件类型不受支持")
	}
	// 校验文件大小
	if fileHeader.Size > (4 * 1024 * 1024) {
		return nil, "", fmt.Errorf("图片大小4MB以内")
	}
	// 打开文件流
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", err
	}
	// 结束时关闭文件流
	defer file.Close()
	// 读取收到的图像到文件流
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, "", err
	}
	// 返回文件流
	return buf.Bytes(), fileHeader.Filename, nil
}

// SaveFileToLocal 保存文件到本地
// @params filePath string 保存路径
// @params data     []byte 数据流
// @return          error  错误信息
func SaveFileToLocal(filePath string, data []byte) error {
	// 取出路径部分
	dir := filepath.Dir(filePath)
	// 创建文件夹
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	// 保存文件
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

// MathPagination 分页计算
// @params total    int64   数据总量
// @params maxNum   *int64  每页最大输出数量
// @params currPage *int64 当前页码
// @return          int64   总页数
// @return          int64   偏移量
// @return          error   错误信息
func MathPagination(total int64, maxNum *int, currPage *int) (totalPages int, startLine int, err error) {
	// 捕获异常
	defer func() {
		if errs := recover(); errs != nil {
			err = fmt.Errorf("分页计算发生错误：%s", errs)
			logs.Error(err.Error())
		}
	}()
	// 限制每页最大数据量
	if *maxNum > 100 {
		*maxNum = 100
	} else if *maxNum <= 0 {
		*maxNum = 30
	}
	// 强制转换
	temTotal := int(total)
	// 计算总页数
	totalPages = temTotal / *maxNum
	if temTotal%*maxNum > 0 {
		totalPages++
	}
	// 限制当前页码
	if *currPage > totalPages {
		*currPage = totalPages
	} else if *currPage <= 0 {
		*currPage = 1
	}
	// 计算偏移查询量
	startLine = (*currPage - 1) * (*maxNum)
	return totalPages, startLine, nil
}

// AesDecrypt AES解密
// @params keyStr     string 密钥
// @params base64Data string Base64之后的AES密文
// @return            string 解密后的数据
// @return            error  错误信息
func AesDecrypt(keyStr string, base64Data string) (string, error) {
	// 密钥转换成流
	key := []byte(keyStr)
	// base64数据解密成流
	aesDataBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(aesDataBytes))
	//执行解密
	blockMode.CryptBlocks(crypted, aesDataBytes)
	//去除填充
	length := len(crypted)
	if length == 0 {
		return "", fmt.Errorf("")
	}
	//获取填充的个数
	unPadding := int(crypted[length-1])
	crypted = crypted[:(length - unPadding)]
	// 返回解密结果
	return string(crypted), nil
}

// ZipDeCompress Zip解压文件
// @params zipFile  string 压缩文件路径
// @params dstPath  string 解压后的文件夹路径
// @return          error  错误信息
func ZipDeCompress(zipFile string, dstPath string) error {
	// 不管三七二十一，先创建目标文件夹
	err := os.MkdirAll(dstPath, 0755)
	if err != nil {
		return err
	}
	// 使用zip打开压缩包
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	// 最外层文件夹的长度
	dirLen := 0
	// 遍历压缩包的内容
	// 注意：reader.File获取到的是压缩包内的所有文件，包括子文件夹下的文件
	for i, file := range reader.File {
		// 忽略掉最外层的压缩包名文件夹
		if i == 0 {
			dirLen = len(file.Name)
			continue
		}
		// 文件夹就不解压出来了
		if file.FileInfo().IsDir() { // 文件夹
			// 不管三七二十一，先创建目标文件夹
			err := os.MkdirAll(dstPath+"/"+file.Name[dirLen:], 0755)
			if err != nil {
				return err
			}
			// 文件夹创建完毕就跳过吧
			continue
		} else { // 文件
			// 打开压缩包内的文件
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()
			// 在文件夹内创建这个文件
			destFile, err := os.Create(dstPath + "/" + file.Name[dirLen:])
			if err != nil {
				return err
			}
			defer destFile.Close()
			// 执行文件拷贝
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}
	// 全部解压完毕
	return nil
}
