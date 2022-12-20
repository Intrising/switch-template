package handlers

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Intrising/intri-core/hal"
	"github.com/Intrising/intri-core/services"

	engineSystem "github.com/Intrising/intri-core/engine/system"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
)

type SystemHandler struct {
	ctx     context.Context
	service commonpb.ServicesEnumTypeOptions

	srv *services.SystemServer
}

func (c *SystemHandler) GetRegisteredMainService() commonpb.ServicesEnumTypeOptions {
	return commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
}

func (c *SystemHandler) getCallBack() *engineSystem.CallBack {
	return &engineSystem.CallBack{
		Ctx:          c.ctx,
		EventClient:  hal.EventClientInit(c.ctx, c.service),
		DeviceClient: hal.DeviceClientInit(c.ctx, c.service),
		// HardwareClient: hal.HardwareClientInit(c.ctx, c.service),
	}
}

func (c *SystemHandler) getDefaultConfig() *systempb.Config {
	return &systempb.Config{
		Identification: &systempb.IdentificationSetting{
			Name:        "test1",
			Description: "test2",
			Location:    "test3",
			Contact:     "test4",
		},
		Oob: &systempb.OOBServicePortSetting{
			IPv4: &systempb.OOBServicePortNetworkSetting{
				Enabled: false,
				IPAddr:  "",
				Netmask: "",
			},
		},
		Watchdog: &systempb.WatchDogSetting{},
		Logout:   &systempb.AutoLogoutSetting{},
	}
}

// func (c *SystemHandler) getConfig() {
// 	c.configClient = hal.ConfigClientInit(c.ctx, c.service)

// 	saveCfg, _ := c.configClient.GetSystemConfig()
// 	defaultCfg := c.defaultConfig()
// 	if saveCfg == nil {
// 		saveCfg = defaultCfg
// 		c.configClient.SetSystemConfig(configpb.ConfigTypeOptions_CONIFG_TYPE_RUNNING, defaultCfg)
// 	}
// 	c.configClient.SetSystemConfig(configpb.ConfigTypeOptions_CONIFG_TYPE_DEFAULT, defaultCfg)
// }

func (c *SystemHandler) Init(ctx context.Context, grpcSrvConn *grpc.Server) {
	c.ctx = ctx
	c.service = commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
	cb := c.getCallBack()
	engineSystem.RegisterCallBack(cb, nil)

	c.srv = &services.SystemServer{
		Cfg:         c.getDefaultConfig(),
		EventClient: cb.EventClient.(*hal.EventClient),
	}

	systempb.RegisterConfigServiceServer(grpcSrvConn, c.srv)
	systempb.RegisterRunServer(grpcSrvConn, c.srv)
	systempb.RegisterInternalServer(grpcSrvConn, c.srv)
}

func init() {
	registerHandler(new(SystemHandler))
}
