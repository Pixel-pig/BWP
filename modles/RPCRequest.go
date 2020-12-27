package modles

type RPCRequest struct {
	Id      int64         `json:"id"`
	Method  string        `json:"method"`
	JsonRpc string        `json:"jsonrpc"`
	Params  []interface{} `json:"params"`
}
