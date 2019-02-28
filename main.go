package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/kataras/iris"
	"github.com/lsj575/wordfilter/bootstrap"
	"github.com/lsj575/wordfilter/routes"
	"os"
	"strconv"
)

const (
	PORT int = 9712
)

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Token敏感词检测系统", "Miracle")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	if len(os.Args) < 2 {
		app.Run(iris.Addr(fmt.Sprintf(":%d", PORT)))
	} else {
		args := os.Args
		port, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("args is wrong", err)
		}
		app.Listen(fmt.Sprintf(":%d", port))
	}
}
