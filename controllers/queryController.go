package controllers

import (
	"BWP/modles"
	"BWP/service"
	"github.com/astaxie/beego"
)

type QueryController struct {
	beego.Controller
}

func (q *QueryController) Get()  {
	q.TplName = "query.html"
}

func (q *QueryController) Post() {
	var query modles.QueryInfo
	err := q.ParseForm(&query)
	if err != nil {
		q.Ctx.WriteString("数据解析失败！请重试")
		return
	}
	_,err = query.SeverBlock()
	if err != nil {
		q.Ctx.WriteString("数据库保存失败！请重试")
		return
	}
	if query.Method == modles.GETBLOCKCOUNT {
		count := service.GetBlockCount()
		q.Data["s"] = "当前节点区块的总数为："
		q.Data["count"] = count
		q.TplName = "query.html"
	}
	if query.Method == modles.GETBESTBLOCKHASH {
		besthash := service.GetBestBlockHash()
		q.Data["s"] = "当前节点最新区块的hash值为："
		q.Data["count"] = besthash
		q.TplName = "query.html"
	}
	if query.Method == modles.GETBLOCK {
		blockInfo := service.GetBlock(query.Param)
		q.Data["s"] = "指定区块的信息："
		q.Data["count"] = blockInfo
		q.TplName = "query.html"
	}
}