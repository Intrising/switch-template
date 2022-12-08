package services

import (
	"context"
	"fmt"

	systempb "github.com/Intrising/intri-type/core/system"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/wrappers"
)

type SysteImplmenetServer struct {
	cfg *systempb.Config
}

func (c *SysteImplmenetServer) Init(ctx context.Context, in *systempb.Config) (*emptypb.Empty, error) {
	fmt.Println("init : ", in)
	// set config to engine if necessary
	// c.cfgSrv.Cfg = &aclpb.Config{}
	// c.cfgSrv.Add_Config_AccessControlList(ctx, &aclpb.AccessControlEntry_List{List: in.AccessControlList})
	c.cfg = proto.Clone(in).(*systempb.Config)
	fmt.Println("Init c.cfg = ", c.cfg)
	return empty, nil
}

func (c *SysteImplmenetServer) Get_Config(ctx context.Context, in *emptypb.Empty) (*systempb.Config, error) {
	fmt.Println("c.cfg = ", c.cfg)
	return proto.Clone(c.cfg).(*systempb.Config), nil
}

func (c *SysteImplmenetServer) Set_Config_Identification_SysName(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	c.cfg.Identification.SysName = in.GetValue()
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Identification_SysLocation(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	c.cfg.Identification.SysLocation = in.GetValue()
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Identification_SysGroup(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	c.cfg.Identification.SysGroup = in.GetValue()
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Identification_SysContact(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	c.cfg.Identification.SysContact = in.GetValue()
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Oob_IPv_4_Enabled(ctx context.Context, in *wrappers.BoolValue) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Oob_IPv_4_IPAddr(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Oob_IPv_4_Netmask(ctx context.Context, in *wrappers.StringValue) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Watchdog_Enabled(ctx context.Context, in *wrappers.BoolValue) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Watchdog_TriggerTime(ctx context.Context, in *wrappers.Int32Value) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Logout_Enabled(ctx context.Context, in *wrappers.BoolValue) (*emptypb.Empty, error) {
	return empty, nil
}

func (c *SysteImplmenetServer) Set_Config_Logout_LogoutTime(ctx context.Context, in *wrappers.Int32Value) (*emptypb.Empty, error) {
	return empty, nil
}
