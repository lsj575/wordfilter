package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/lsj575/wordfilter/datasource"
	"github.com/lsj575/wordfilter/models"
	"log"
	"os"
	"strings"
	"sync"
)

type SensitiveWordController struct {
	Ctx iris.Context
}

func (c *SensitiveWordController) Post() {
	words := c.Ctx.FormValue("words")
	sensitiveTrie := datasource.InstanceSensitiveWord()
	wordArr := strings.Split(words, ",")
	errCode := 0
	msg := "OK"
	for _, word := range wordArr {
		if !sensitiveTrie.Find(word) {
			err := c.traceFile(word)
			if err != nil {
				log.Println("SensitiveWordController Post traceFile ", err.Error())
				errCode = 1
				msg = fmt.Sprintf("SensitiveWordController Post traceFile %T", err.Error())
				break
			}
			sensitiveTrie.Add(word)
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

func (c *SensitiveWordController) traceFile(str string) error {
	locker := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()
	fd, err := os.OpenFile(datasource.SensitiveFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	str += "\n"
	buf:=[]byte(str)
	fd.Write(buf)
	fd.Close()
	return nil
}
