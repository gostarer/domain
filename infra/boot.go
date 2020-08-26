package infra

import (
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

type BootApplication struct {
	IsTest    bool
	conf      kvs.ConfigSource
	goStarCtx GoStarerContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	boot := &BootApplication{
		conf:      conf,
		goStarCtx: GoStarerContext{},
	}
	boot.goStarCtx.SetProps(conf)
	return boot
}

func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

func (b *BootApplication) init() {
	logrus.Info("init gostarers")
	for _, v := range GetGoStars() {
		typ := reflect.TypeOf(v)
		logrus.Debugf("init:", typ.String())
		v.Init(b.goStarCtx)
	}
}

func (b *BootApplication) setup() {
	logrus.Info("all star setup")
	for _, v := range GetGoStars() {
		typ := reflect.TypeOf(v)
		logrus.Debugf("setup:", typ.String())
		v.Setup(b.goStarCtx)
	}
}

func (b *BootApplication) start() {
	logrus.Info("all star start")
	for i, v := range GetGoStars() {
		typ := reflect.TypeOf(v)
		logrus.Debugf("start:", typ.String())
		if v.StartBlocking() {
			if i+1 == len(GetGoStars()) {
				v.Start(b.goStarCtx)
			} else {
				logrus.Infof("blocking:%v", typ.String())
				go v.Start(b.goStarCtx)
			}
		} else {
			v.Start(b.goStarCtx)
		}
	}
}

func (b *BootApplication) stop() {
	logrus.Info("all star stop")
	for _, v := range GetGoStars() {
		typ := reflect.TypeOf(v)
		logrus.Debugf("start:", typ.String())
		v.Stop(b.goStarCtx)
	}
}
