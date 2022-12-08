package services

import (
	"context"
	"sync"

	systempb "github.com/Intrising/intri-type/core/system"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SystemServer struct {
	mtx sync.RWMutex
	systempb.PreImplementConfigurationServiceServer
	systempb.UnimplementedInternalServicesServer
	systempb.UnimplementedRunServicesServer
	PreImplementServer *SysteImplmenetServer
}

func (c *SystemServer) Run_Get_Info(context.Context, *emptypb.Empty) (*systempb.Info, error) {
	return &systempb.Info{
		Oid: "1232189312893",
	}, nil
}

func (c *SystemServer) Run_Get_Security_Info(context.Context, *emptypb.Empty) (*systempb.SecurityInfo, error) {
	return &systempb.SecurityInfo{
		SoftWareKernelVersion: "1.1.1",
	}, nil
}
