package modles

type RequestResult struct {
	Id     int64       `json:"id"`
	Error  ResultError `json:"error"`
	Result interface{} `json:"result"`
}

type ResultError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
