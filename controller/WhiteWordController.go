package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/lsj575/wordfilter/datasource"
	"github.com/lsj575/wordfilter/models"
	"log"
	"os"
	"strings"
)

type WhiteWordController struct {
	Ctx iris.Context
}

func (c *WhiteWordController) Post() {
	words := c.Ctx.FormValue("words")
	whiteTrie := datasource.InstanceWhiteWord()
	wordArr := strings.Split(words, ",")
	errCode := 0
	msg := "OK"
	for _, word := range wordArr {
		if !whiteTrie.Find(word) {
			err := c.traceFile(word)
			if err != nil {
				log.Println("WhiteWordController Post traceFile ", err.Error())
				errCode = 1
				msg = fmt.Sprintf("WhiteWordController Post traceFile %T", err.Error())
				break
			}
			whiteTrie.Add(word)
		} else {
			errCode = 1
			msg = fmt.Sprintf("%s is exist", word)
		}
	}
	data := models.ApiReturn{
		ErrCode: errCode,
		Msg:     msg,
		Data:    nil,
	}
	c.Ctx.JSON(data)
}

func (c *WhiteWordController) traceFile(str string) error {
	fd, err := os.OpenFile(datasource.WhiteFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	str += "\n"
	buf:=[]byte(str)
	fd.Write(buf)
	fd.Close()
	return nil
}
