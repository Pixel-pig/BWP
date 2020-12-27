package service

import (
	"BWP/modles"
	"BitcoinRPC/entity"
	"BitcoinRPC/utils"
	"fmt"
)

// 获取当前节点区块的总数
func GetBlockCount() float64{
	json := utils.PrepareJson(entity.GETBLOCKCOUNT)
	response, body := utils.HttpRequest(json)
	//状态码
	if response.StatusCode == 200 {
		result := utils.JosnDecoding(body)
		return result.Result.(float64)
	} else {
		fmt.Println("请求失败！，状态码为：", response.StatusCode)
		return -1
	}
}

// 获取当前节点最新区块的hash值
func GetBestBlockHash() string {
	json := utils.PrepareJson(entity.GETBESTBLOCKHASH)
	response, body := utils.HttpRequest(json)
	//状态码
	if response.StatusCode == 200 {
		result := utils.JosnDecoding(body)
		return result.Result.(string)
	} else {
		fmt.Println("请求失败！，状态码为：", response.StatusCode)
		return ""
	}
}

// 根据指定的hash值获取指定的区块信息
func GetBlock(hash string)  *modles.BlockInfo {
	json := utils.PrepareJson(entity.GETBLOCK,hash)
	response, body := utils.HttpRequest(json)
	//状态码
	if response.StatusCode == 200 {
		result := utils.JosnDecoding(body)
		var blockInfo modles.BlockInfo
		block := result.Result.(map[string]interface{})
		blockInfo.Hash = block["hash"].(string)
		blockInfo.Size = block["size"].(float64)
		blockInfo.Height = block["height"].(float64)
		return &blockInfo
	} else {
		fmt.Println("请求失败！，状态码为：", response.StatusCode)
		return nil
	}
}
