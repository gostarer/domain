package base

import (
	"github.com/gostarer/domain/infra"
	"net"
	"net/rpc"
	"reflect"

	log "github.com/sirupsen/logrus"
)

var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	Check(rpcServer)
	return rpcServer
}

func RpcRegister(ri interface{}) {
	typ := reflect.TypeOf(ri)
	log.Infof("rpc register :%s", typ.String())
	RpcServer().Register(ri)
}

type GoRPCStarer struct {
	infra.BaseGoStarer
	server *rpc.Server
}

func (s *GoRPCStarer) Init(ctx infra.GoStarerContext) {
	s.server = rpc.NewServer()
	rpcServer = s.server
}

func (s *GoRPCStarer) Start(ctx infra.GoStarerContext) {
	port := ctx.Props().GetDefault("app.rpc.port", "8082")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	log.Info("tcp port listened for rpc:", port)
	go s.server.Accept(listener)
}
