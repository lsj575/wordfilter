package middleware

import (
	"encoding/base64"
	"github.com/kataras/iris"
	"github.com/lsj575/wordfilter/models"
	"strconv"
	"strings"
	"time"
)

var CheckSign = func(ctx iris.Context) {
	sign := ctx.GetHeader("sign")
	decodeSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		data := &models.ApiReturn{
			ErrCode: 1,
			Msg:     "sign is wrong",
			Data:    nil,
		}
		ctx.JSON(data)
		ctx.StopExecution()
	}
	arrSign := strings.Split(string(decodeSign), "%7C%7C")
	if len(arrSign) < 3 {
		data := &models.ApiReturn{
			ErrCode: 1,
			Msg:     "sign is wrong",
			Data:    nil,
		}
		ctx.JSON(data)
		ctx.StopExecution()
	}
	signTime, err := strconv.Atoi(arrSign[1])
	if err != nil {
		data := &models.ApiReturn{
			ErrCode: 1,
			Msg:     "sign is wrong",
			Data:    nil,
		}
		ctx.JSON(data)
		ctx.StopExecution()
	}
	timeNow := int(time.Now().Unix())
	if arrSign[0] != "sensitive" ||
		(timeNow < signTime || signTime < timeNow - 10) ||
		arrSign[2] != "token" {
		data := &models.ApiReturn{
			ErrCode: 1,
			Msg:     "sign is wrong",
			Data:    nil,
		}
		ctx.JSON(data)
		ctx.StopExecution()
	}
	ctx.Next()
}
