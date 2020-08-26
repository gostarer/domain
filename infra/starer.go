package infra

import (
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

const (
	KeyProps = "_conf"
)

type GoStarerContext map[string]interface{}

func (gs GoStarerContext) Props() kvs.ConfigSource {
	p := gs[KeyProps]
	if p == nil {
		panic("配置没有被初始化")
	}
	return p.(kvs.ConfigSource)
}

func (gs GoStarerContext) SetProps(conf kvs.ConfigSource) {
	gs[KeyProps] = conf
}

type GoStarer interface {
	Init(GoStarerContext)
	Setup(GoStarerContext)
	Start(GoStarerContext)
	StartBlocking() bool
	Stop(GoStarerContext)
	//PriorityGroup() PriorityGroup
	//Priority() int
}

type goStarRegister struct {
	nonBlockingGoStars []GoStarer
	blockingGoStars    []GoStarer
}

func (gs *goStarRegister) Register(goStar GoStarer) {
	if goStar.StartBlocking() {
		gs.blockingGoStars = append(gs.blockingGoStars, goStar)
	} else {
		gs.nonBlockingGoStars = append(gs.nonBlockingGoStars, goStar)
	}
	typ := reflect.TypeOf(goStar)
	logrus.Infof("register service:%s", typ.String())
}

func (gs *goStarRegister) AllGoStarer() []GoStarer {
	gostarers := make([]GoStarer, 0)
	gostarers = append(gostarers, gs.nonBlockingGoStars...)
	gostarers = append(gostarers, gs.blockingGoStars...)
	logrus.Infof("AllGoStarer:%v", gostarers)
	return gostarers
}

var GoStarRegister *goStarRegister = &goStarRegister{}

func GetGoStars() []GoStarer {
	return GoStarRegister.AllGoStarer()
}

func Register(starer GoStarer) {
	GoStarRegister.Register(starer)
}

type BaseGoStarer struct {
}

func (s *BaseGoStarer) Init(ctx GoStarerContext)  {}
func (s *BaseGoStarer) Setup(ctx GoStarerContext) {}
func (s *BaseGoStarer) Start(ctx GoStarerContext) {}
func (s *BaseGoStarer) Stop(ctx GoStarerContext)  {}
func (s *BaseGoStarer) StartBlocking() bool       { return false }

//func (s *BaseGoStarer) PriorityGroup() PriorityGroup { return BaseResourceGroup }
//func (s *BaseGoStarer) Priority() int                { return DEFAULT_PRIORITY }
