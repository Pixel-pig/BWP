package rpcUtils

import (
	"BWPjy/models"
	"net/http"
)

//获取最新区块的hash
func GetBestBlockHash() interface{} {
	reqbytes := PrepareJSON(models.GETBESTBLOCKHASH)
	response, body := DoPost(reqbytes)
	if response == nil {
		return nil
	}
	if response.StatusCode == http.StatusOK {
		count := Analysis(response, body)
		return count.Result
	}
	return nil
}

//该方法用于获取比特币当前节点的区块数
func GetBlockCount() interface{} {
	reqbytes := PrepareJSON(models.GETBLOCKCOUNT)
	response, body := DoPost(reqbytes)
	if response == nil {
		return nil
	}
	if response.StatusCode == http.StatusOK {
		count := Analysis(response, body)
		return count.Result
	}
	return nil
}

//根据区块hash获取区块信息
func GetBlock(params interface{}) interface{} {
	reqbytes := PrepareJSON(models.GETBLOCK, params)
	//fmt.Println(reqbytes)
	response, body := DoPost(reqbytes)
	//fmt.Println(response.StatusCode)
	if response == nil {
		return nil
	}
	if response.StatusCode == http.StatusOK {
		data := Analysis(response, body)
		return data.Result
	}
	return nil
}
