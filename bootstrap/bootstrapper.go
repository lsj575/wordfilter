package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lsj575/lottery/config"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName string
	AppOwner string
	AppSpawDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawDate: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true)
	htmlEngine.AddFunc("FromUnixTimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixTime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(config.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(context iris.Context) {
		err := iris.Map{
			"app": b.AppName,
			"status": context.GetStatusCode(),
			"message": context.Values().GetString("message"),
		}
		if jsonOutput := context.URLParamExists("json"); jsonOutput {
			context.JSON(err)
			return
		}
		context.ViewData("Err", err)
		context.ViewData("Title", "Error")
		context.View("shared/error.html")
	})
}

func (b *Bootstrapper) Configure(cfgs ...func(b *Bootstrapper)) {
	for _, cfg := range cfgs {
		cfg(b)
	}
}

func (b *Bootstrapper) setupCron() {
	// TODO
}

const (
	StaticAssets = "./public/"
	Favicon = "favicon.ico"
)

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//b.SetupViews("./views")
	//b.SetupErrorHandlers()
	//b.Favicon(StaticAssets + Favicon)
	//b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)
	//b.setupCron()

	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
