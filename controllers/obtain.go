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

var count interface{}
var getBast interface{}

func (h *Obtain) Post() {
	var users models.Users
	err := h.ParseForm(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch users.Method {
	case entlity.GETBLOCKCOUNT :
		count = util.GetBlockCount()
		//temp = count.(string)
		h.Data["s"]="区块的总数:"
		h.Data["count"]=count
	case entlity.GETBESTBLOCKHASH:
		getBast = util.GetBestBlockHash()
		h.Data["s"]="区块的哈希:"
		h.Data["getBast"]=getBast
	default:
		fmt.Println("方法错误，请查看手册")
	}
	//if users.Method == entlity.GETBLOCKCOUNT {
	//	count = util.GetBlockCount()
	//	//temp = count.(string)
	//	h.Data["s"]="区块的总数:"
	//	h.Data["count"]=count
	//	return
	//} else {
	//	h.Ctx.WriteString("方法错误，请查看手册")
	//}
	h.TplName = "obtain.html"
}
