package entlity

type Blockinfo struct {
	Headers              float64
	Chain                string
	SizeOnDisk           float64
	Mediantime           float64
	Blocks               float64
	Pruned               bool
	Warnings             string
	Difficulty           float64
	Chainwork            string
	Verificationprogress float64
	Bestblockhash        string
	Initialblockdownload bool
}
