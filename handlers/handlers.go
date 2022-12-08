package handlers

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	commonpb "github.com/Intrising/intri-type/common"
	utilsLog "github.com/Intrising/intri-utils/log"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
	"google.golang.org/grpc"
)

func GetRuntimeFrame(skipIndex int) runtime.Frame {
	p := make([]uintptr, 1)
	n := runtime.Callers(skipIndex+2, p)
	if n > 0 {
		cframes := runtime.CallersFrames(p[:n])
		cf, _ := cframes.Next()
		return cf
	}
	return runtime.Frame{File: "unknown.file", Function: "unknown.func", Line: -1}
}

var (
	registerHandlerMapping = make(map[commonpb.ServicesEnumTypeOptions][]IHandler)
)

func registerHandler(h IHandler) {
	if _, ok := registerHandlerMapping[h.GetRegisteredMainService()]; !ok {
		registerHandlerMapping[h.GetRegisteredMainService()] = make([]IHandler, 0)
	}
	registerHandlerMapping[h.GetRegisteredMainService()] = append(registerHandlerMapping[h.GetRegisteredMainService()], h)
}

type IHandler interface {
	GetRegisteredMainService() commonpb.ServicesEnumTypeOptions
	Init(ctx context.Context, grpcSrvConn *grpc.Server)
}

type Handler struct{}

func (c *Handler) debugPanic(parent context.Context, r interface{}) {
	traceList := make([]string, 0)

	for i := 0; i < 100; i++ {
		frame := GetRuntimeFrame(i)
		if frame.Line == -1 {
			break
		}
		traceList = append(traceList, fmt.Sprintf("[recover] trace: %04d file(%s): %s:%d",
			i,
			frame.Function,
			frame.File,
			frame.Line,
		))
	}

	utilsLog.Info("[recover] panic: %#v\n%s", r, strings.Join(traceList, "\n"))
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
