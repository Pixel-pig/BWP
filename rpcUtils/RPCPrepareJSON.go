package rpcUtils

import (
	"BWPjy/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//rpc通信的json格式的请求数据
func PrepareJSON(method string, params ...interface{}) string {
	rpcRequest := models.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: models.RPCVERSION,
		Params:params,
	}
	//if params != nil {
	//	rpcRequest.Params = params
	//}
	reqBytes, err := json.Marshal(&rpcRequest)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(reqBytes)
}

//请求的响应结果
func DoPost(Body string) (*http.Response, []byte) {
	//1发起一个post请求
	client := &http.Client{}
	request, err := http.NewRequest("POST", models.RPCURL, strings.NewReader(string(Body)))
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	//2请求头设置
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	authStr := models.RPCUser + ":" + models.RPCPASSWORD
	authBase64 := Base64Encode([]byte(authStr))
	request.Header.Add("Authorization", "Basic "+string(authBase64))
	//3返回响应类：
	response, err := client.Do(request)
	//fmt.Println(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	return response, body
}

//解析数据
func Analysis(response *http.Response, body []byte) *models.Result {
	if response.StatusCode == http.StatusOK {
		result := models.Result{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println("请求成功")
		return &result
	} else {
		fmt.Println("请求失败")
	}
	return nil
}
