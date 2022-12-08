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
		// ConfigClient:   hal.ConfigClientInit(c.ctx, c.service),
		// HardwareClient: hal.HardwareClientInit(c.ctx, c.service),
	}
}

// func (c *SystemHandler) defaultConfig() *systempb.Config {
// 	return &systempb.Config{
// 		SysName:     "os5-switch",
// 		SysLocation: "",
// 		SysGroup:    "",
// 		SysContact:  "",
// 	}
// }

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
	// engineSystem.RegisterCallBack(c.getCallBack())

	c.srv = &services.SystemServer{}
	c.srv.PreImplementServer = &services.SysteImplmenetServer{}
	c.srv.NeedImpInit = c.srv.PreImplementServer
	c.srv.NeedImpCfg = c.srv.PreImplementServer

	systempb.RegisterConfigurationServiceServer(grpcSrvConn, c.srv)
	systempb.RegisterRunServicesServer(grpcSrvConn, c.srv)
	systempb.RegisterInternalServicesServer(grpcSrvConn, c.srv)
}

func init() {
	registerHandler(new(SystemHandler))
}
