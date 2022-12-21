package handlers

import (
	"context"
	"reflect"

	commonpb "github.com/Intrising/intri-type/common"
	utilsLog "github.com/Intrising/intri-utils/log"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
	"google.golang.org/grpc"
)

var (
	registerHandlerMapping = make(map[commonpb.ServicesEnumTypeOptions][]IHandler)
)

type IHandler interface {
	GetRegisteredMainService() commonpb.ServicesEnumTypeOptions
	Init(ctx context.Context, grpcSrvConn *grpc.Server)
}

type Handler struct{}

func registerHandler(h IHandler) {
	if _, ok := registerHandlerMapping[h.GetRegisteredMainService()]; !ok {
		registerHandlerMapping[h.GetRegisteredMainService()] = make([]IHandler, 0)
	}
	registerHandlerMapping[h.GetRegisteredMainService()] = append(registerHandlerMapping[h.GetRegisteredMainService()], h)
}

func (c *Handler) handleInit(parent context.Context, service commonpb.ServicesEnumTypeOptions, hs []IHandler) {
	utilsLog.Info("Init Group '%s' Handlers", service)
	srvConn := utilsRpc.NewServerConn(parent, service)
	grpcSrvConn := srvConn.GetGrpcServer()
	for _, h := range hs {
		utilsLog.Info("Init Handler %s", reflect.TypeOf(h))
		h.Init(parent, grpcSrvConn)
	}
	go srvConn.Run()
}

func (c *Handler) Run(parent context.Context) {
	for service, hs := range registerHandlerMapping {
		c.handleInit(parent, service, hs)
	}
}
