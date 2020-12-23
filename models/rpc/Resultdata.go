package rpc

//比特币返回的结果
type Resultdata struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
	Id     int         `json:"id"`
}
