package base

import (
	"github.com/gostarer/domain/infra"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type HookStarer struct {
	infra.BaseGoStarer
}

func (s *HookStarer) Init(ctx infra.GoStarerContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			log.Info("notify:", c)
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()
}

func (s *HookStarer) Start(ctx infra.GoStarerContext) {
	starers := infra.GetGoStars()
	for _, s := range starers {
		typ := reflect.TypeOf(s)
		log.Infof("[Register notify stop]:%s.stop()", typ.String())
		Register(func() {
			s.Stop(ctx)
		})
	}
}
