package controllers

import "github.com/astaxie/beego"

type QueryController struct {
	beego.Controller
}

func (q *QueryController) Get()  {
	q.TplName = "query.html"
}

func (q *QueryController) Post() {

}