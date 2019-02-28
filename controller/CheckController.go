package controller

import (
	"github.com/kataras/iris"
	"github.com/lsj575/wordfilter/datasource"
	"github.com/lsj575/wordfilter/models"
)

type CheckController struct {
	Ctx iris.Context

}

func (c *CheckController) Post() {
	var ret int
	content := c.Ctx.FormValue("content")
	sensitiveTrie := datasource.InstanceSensitiveWord()
	result, find := sensitiveTrie.Replace(content, datasource.InstanceWhiteWord())
	if len(find) > 0 {
		ret = 1
	} else {
		ret = 0
	}
	data := models.ApiReturn{
		ErrCode: 0,
		Msg:     "OK",
		Data:    models.CheckResult{
			Ret:    ret,
			Result: result,
			Find:   find,
		},
	}
	c.Ctx.JSON(data)
}

func (c *CheckController) Get() {
	var ret int
	content := c.Ctx.URLParamDefault("content", "")
	sensitiveTrie := datasource.InstanceSensitiveWord()
	result, find := sensitiveTrie.Replace(content, datasource.InstanceWhiteWord())
	if len(find) > 0 {
		ret = 1
	} else {
		ret = 0
	}
	data := models.ApiReturn{
		ErrCode: 0,
		Msg:     "OK",
		Data:    models.CheckResult{
			Ret:    ret,
			Result: result,
			Find:   find,
		},
	}
	c.Ctx.JSON(data)
}