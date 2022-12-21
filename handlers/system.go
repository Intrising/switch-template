package handlers

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/Intrising/intri-core/hal"
	"github.com/Intrising/intri-core/services"

	engineSystem "github.com/Intrising/intri-core/engine/system"

	utilsLog "github.com/Intrising/intri-utils/log"
	utilsMisc "github.com/Intrising/intri-utils/misc"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"
)

type SystemHandler struct {
	initReady bool
	ctx       context.Context
	service   commonpb.ServicesEnumTypeOptions

	cfg *systempb.Config
	srv *services.SystemServer
	cb  *engineSystem.CallBack
}

func (c *SystemHandler) GetRegisteredMainService() commonpb.ServicesEnumTypeOptions {
	return commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
}

func (c *SystemHandler) getCallBack() {
	unions := []*eventpb.InternalTypeUnionEntry{
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_SYSTEM},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_NTP},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_BOOT},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_BUTTON},
		// init event
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE, ServicesType: c.service},
		// save config event
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_CONFIG, ServicesType: commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CONFIG},
	}

	required := []commonpb.ServicesEnumTypeOptions{}

	c.cb = &engineSystem.CallBack{
		Ctx:          c.ctx,
		EventClient:  hal.EventClientInit(c.ctx, c.service, unions, required),
		DeviceClient: hal.DeviceClientInit(c.ctx, c.service),
		// HardwareClient: hal.HardwareClientInit(c.ctx, c.service),
	}
}

func (c *SystemHandler) getDefaultConfig() *systempb.Config {
	name := c.cb.DeviceClient.GetDeviceInfo().Model
	description := c.cb.DeviceClient.GetBoardInfo().SystemDescription
	return &systempb.Config{
		Identification: &systempb.IdentificationSetting{
			Name:        name,
			Description: description,
			Location:    "Taiwan Taipei",
			Contact:     "",
		},
		Oob: &systempb.OOBServicePortSetting{
			IPv4: &systempb.OOBServicePortNetworkSetting{
				Enabled: false,
				IPAddr:  "192.168.0.1",
				Netmask: "255.255.255.0",
			},
		},
		Watchdog: &systempb.WatchDogSetting{
			Enabled:     false,
			TriggerTime: 60,
		},
		Logout: &systempb.AutoLogoutSetting{
			Enabled:    false,
			LogoutTime: 60,
		},
	}
}

func (c *SystemHandler) saveConfig(in *systempb.Config, path string) {
	fmt.Println("saveConfig : enter = ", path, in)
	filePath := path + c.service.String() + ".yml"
	cfg := proto.Clone(in).(*systempb.Config)
	utilsMisc.SaveProtoMessageToFile(cfg, filePath)
}

func (c *SystemHandler) loadSaveConfig(path string) (*systempb.Config, error) {
	filePath := path + c.service.String() + ".yml"
	out := &systempb.Config{}
	err := utilsMisc.LoadProtoMessageFromFile(out, filePath)
	return out, err
}

func (c *SystemHandler) getConfig() {
	defaultCfg := c.getDefaultConfig()
	path := c.cb.DeviceClient.GetPath()

	// save default config
	c.saveConfig(defaultCfg, path.GetConfigDefault())

	var err error
	// load save config
	c.cfg, err = c.loadSaveConfig(path.GetConfigSaved())
	if err != nil {
		// save default config to saved
		c.saveConfig(defaultCfg, path.GetConfigSaved())
		c.cfg = proto.Clone(defaultCfg).(*systempb.Config)
	}
}

func (c *SystemHandler) initConfig() {
	if c.initReady {
		utilsLog.Info("init conifg ready")
		return
	}
	c.initReady = true

	c.getConfig()
	engineSystem.RegisterCallBack(c.cb, c.cfg)
	c.srv.InitConfig(c.cfg)
}

func (c *SystemHandler) listenEvent() {
	utilsLog.Info("listenEvent : enter")
	for {
		evt, err := c.cb.EventClient.ReceiveEvent()
		fmt.Println(evt, err)
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		utilsLog.Info("evt = ", evt)
		switch evt.GetType() {
		case eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE:
			c.initConfig()
		case eventpb.InternalTypeOptions_INTERNAL_TYPE_CONFIG:
			if evt.GetConfig().GetActionOption() == eventpb.ConfigActionTypeOptions_CONFIG_ACTION_TYPE_CONFIG_SAVE {
				c.saveConfig(c.cfg, c.cb.DeviceClient.GetPath().GetConfigSaved())
			}
		}
	}
}

func (c *SystemHandler) Init(ctx context.Context, grpcSrvConn *grpc.Server) {
	c.ctx = ctx
	c.service = c.GetRegisteredMainService()
	c.getCallBack()

	c.srv = &services.SystemServer{
		EventClient: c.cb.EventClient.(*hal.EventClient),
	}

	systempb.RegisterConfigServiceServer(grpcSrvConn, c.srv)
	systempb.RegisterRunServiceServer(grpcSrvConn, c.srv)
	systempb.RegisterInternalServiceServer(grpcSrvConn, c.srv)

	go c.listenEvent()
}

func init() {
	registerHandler(new(SystemHandler))
}
