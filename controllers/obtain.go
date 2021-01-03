package controllers

import (
	"BTCproject/entlity"
	"BTCproject/models"
	"BTCproject/util"
	"fmt"
	"github.com/astaxie/beego"
)

type Obtain struct {
	beego.Controller
}

func (o *Obtain) Get() {
	o.TplName = "obtain.html"
}
func (o *Obtain) Post() {
	var users models.Users
	err := o.ParseForm(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch users.Method {
	case entlity.GETBLOCKCOUNT:
		count := util.GetBlockCount()
		//temp = count.(string)
		o.Data["s"] = "区块的总数:"
		o.Data["count"] = count
	case entlity.GETBESTBLOCKHASH:
		getBast := util.GetBestBlockHash()
		o.Data["s"] = "最新区块的哈希:"
		o.Data["getBast"] = getBast
	case entlity.GETBLOCK:
		information := util.GetBlock(users.Parameter)
		o.Data["s"] = "区块的信息:"
		o.Data["information"] = information
	//case entlity.GETBLOCKHASH:
	//	hash := util.GetBlockHash(users.Parameter)
	//	o.Data["s"] = "指定区块的hash:"
	//	o.Data["hash"] = hash
	case entlity.GETBLOCKCHAININFO:
		blockinfo := util.GetBlockChainInfo()
		o.Data["s"] = "区块链信息:"
		o.Data["blockinfo"] = blockinfo
	case entlity.GETDIFFICULTY:
		difficulty:=util.Getdifficulty()
		o.Data["s"]="区块的难度:"
		o.Data["difficulty"]=difficulty
	default:
		fmt.Println("方法错误，请查看手册")
	}
	o.TplName = "obtain.html"
}
