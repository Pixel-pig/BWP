package utils

import (
	"BitcoinRPC/entity"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpRequest(json []byte) (*http.Response,[]byte) {
	//创建一个客户端
	client := &http.Client{}
	//发起post请求
	request, err := http.NewRequest("POST", entity.RPCURL, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err.Error())
		return nil,nil
	}
	//请求头设置
	authStr := entity.RPCUSER + ":" + entity.RPCPASSWORD
	authBase64 := Base64EnConding(authStr)
	request.Header.Add("Authorization", "Basic "+ authBase64)
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	//读取响应数据
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
		return nil,nil
	}
	body, _ := ioutil.ReadAll(response.Body)

	return response,body

}
