package utils

import (
	"BitcoinRPC/entity"
	"encoding/json"
	"fmt"
	"time"
)

/**
 * 该方法用于准备一个json格式的数据
 */
func PrepareJson(method string,params ...interface{}) []byte {
	rpcRequest := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		JsonRpc: entity.RPCVERSION,
	}
	if params != nil{
		rpcRequest.Params = params
	}
	reqBytes, err := json.Marshal(&rpcRequest)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return reqBytes
}

/**
 * 该函数用于对数据进行反序列化
 */
func JosnDecoding(body []byte) *entity.RequestResult {
	result := entity.RequestResult{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &result
}