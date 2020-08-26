package base

import (
	"github.com/gostarer/domain/infra"

	"github.com/sirupsen/logrus"

	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	Check(props)
	return props
}

type PropsStarer struct {
	infra.BaseGoStarer
}

func (p *PropsStarer) Init(ctx infra.GoStarerContext) {
	props = ctx.Props()
	logrus.Info("init config")
	//GetSystemAccount()
}
