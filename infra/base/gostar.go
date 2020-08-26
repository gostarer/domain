package base

import (
	"github.com/gostarer/domain/infra"

	"github.com/gostarer/web"
)

var gostarWeb *web.Engine

func GoStar() *web.Engine {
	Check(gostarWeb)
	return gostarWeb
}

type GoStarWebServerStarer struct {
	infra.BaseGoStarer
}

func (gwss *GoStarWebServerStarer) Init(ctx infra.GoStarerContext) {
	gostarWeb = initGoStarWeb()
}

func (gwss *GoStarWebServerStarer) Setup(ctx infra.GoStarerContext) {

}

func (gwss *GoStarWebServerStarer) Start(ctx infra.GoStarerContext) {
	port := ctx.Props().GetDefault("app.server.port", "18081")
	gostarWeb.Run(":" + port)
}

func (gwss *GoStarWebServerStarer) StartBlocking() bool {
	return true
}

func initGoStarWeb() *web.Engine {
	gostarWeb := web.New()
	gostarWeb.Use(web.Logger(), web.Recovery())
	return gostarWeb
}
