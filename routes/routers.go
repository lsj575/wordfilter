package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/lsj575/wordfilter/bootstrap"
	"github.com/lsj575/wordfilter/controller"
	"github.com/lsj575/wordfilter/middleware"
)

func Configure(b *bootstrap.Bootstrapper) {
	check := mvc.New(b.Party("/check"))
	check.Router.Use(middleware.CheckSign)
	check.Handle(new(controller.CheckController))

	whiteWord := mvc.New(b.Party("/whiteword"))
	whiteWord.Router.Use(middleware.BasicAuth)
	whiteWord.Handle(new(controller.WhiteWordController))

	sensitiveWord := mvc.New(b.Party("/sensitiveword"))
	sensitiveWord.Router.Use(middleware.BasicAuth)
	sensitiveWord.Handle(new(controller.SensitiveWordController))
}
