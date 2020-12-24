package entlity


type Block struct {
	Bits          string `json:"bits"`
	Hash          string `json:"hash"`
	Confirmations int    `json:"confirmations"`
	Strippedsize  int    `json:"strippedsize"`
	Size          int    `json:"size"`
	Weight        int    `json:"weight"`
	Height        int    `json:"height"`
	Version       int    `json:"version"`
	VersionHex    string `json:"version_hex"`
	Merkleroot    string `json:"merkleroot"`
	//	Tx            [...]string{}
	Time          int64  `json:"time"`
	Mediantime    int64  `json:"mediantime"`
	Nonce         int64  `json:"nonce"`
	Difficulty    int    `json:"difficulty"`
	Chainwork     string `json:"chainwork"`
	NTx           int    `json:"nTx"`
	Nextblockhash string `json:"nextblockhash"`
}
