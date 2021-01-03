package entlity

type Block struct {
	Bits          string  `json:"bits"`
	Hash          string  `json:"hash"`
	Confirmations float64 `json:"confirmations"`
	Strippedsize  float64 `json:"strippedsize"`
	Size          float64 `json:"size"`
	Weight        float64 `json:"weight"`
	Height        float64 `json:"height"`
	Version       float64 `json:"version"`
	VersionHex    string  `json:"version_hex"`
	Merkleroot    string  `json:"merkleroot"`
	Time          float64 `json:"time"`
	Mediantime    float64 `json:"mediantime"`
	Nonce         float64 `json:"nonce"`
	Difficulty    float64 `json:"difficulty"`
	Chainwork     string  `json:"chainwork"`
	NTx           float64 `json:"nTx"`
	Nextblockhash string  `json:"nextblockhash"`
}
