package services

import (
	"context"
	"sync"

	systempb "github.com/Intrising/intri-type/core/system"

	utilsLog "github.com/Intrising/intri-utils/log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SystemServer struct {
	mtx sync.RWMutex
	systempb.UnimplementedConfigurationServiceServer
	systempb.UnimplementedInternalServicesServer
	systempb.UnimplementedRunServicesServer
}

func (c *SystemServer) Init(ctx context.Context, in *systempb.Config) (*emptypb.Empty, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if in == nil {
		utilsLog.Warning("the in == nil")
		return empty, nil
	}

	// set config to engine if necessary
	// c.cfgSrv.Cfg = &aclpb.Config{}
	// c.cfgSrv.Add_Config_AccessControlList(ctx, &aclpb.AccessControlEntry_List{List: in.AccessControlList})
	return empty, nil
}

func (c *SystemServer) Get_Config(context.Context, *emptypb.Empty) (*systempb.Config, error) {
	return &systempb.Config{
		Logout: &systempb.AutoLogoutSetting{
			LogoutTime: 30,
		},
	}, nil
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
