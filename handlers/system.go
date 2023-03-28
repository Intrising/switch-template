package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Intrising/switch-template/hal"
	"github.com/Intrising/switch-template/services"

	engineSystem "github.com/Intrising/switch-template/engine/system"

	utilsLog "github.com/Intrising/intri-utils/log"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"
)

type SystemHandler struct {
	initReady bool
	ctx       context.Context
	service   commonpb.ServicesEnumTypeOptions

	configClient *hal.ConfigClient
	eventClient  *hal.EventClient
	deviceClient *hal.DeviceClient

	cfg *systempb.Config
	srv *SystemProtectServiceServer
	cb  *engineSystem.CallBack
}

type SystemProtectServiceServer struct {
	systempb.ConfigServiceServer
	systempb.RunServiceServer
}

func (c *SystemHandler) registerCallBack() {
	c.cb = &engineSystem.CallBack{
		Ctx:          c.ctx,
		EventClient:  c.eventClient,
		DeviceClient: c.deviceClient,
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

func (c *SystemHandler) saveConfig(in *systempb.Config) {
	c.configClient.SaveSavedConfig(in)
}

func (c *SystemHandler) saveDefaultConfig(in *systempb.Config) {
	cfg := proto.Clone(in).(*systempb.Config)
	c.configClient.SaveDefaultConfig(cfg)
}

func (c *SystemHandler) loadSaveConfig() (*systempb.Config, error) {
	out := &systempb.Config{}
	err := c.configClient.LoadSavedConfig(out)
	return out, err
}

func (c *SystemHandler) getConfig() {
	defaultCfg := c.getDefaultConfig()

	// save default config
	c.saveDefaultConfig(defaultCfg)

	var err error
	// load save config
	c.cfg, err = c.loadSaveConfig()
	if err != nil {
		log.Println("err = ", err)
		// save default config to saved
		c.saveConfig(defaultCfg)
		c.cfg = proto.Clone(defaultCfg).(*systempb.Config)
	}
}

func (c *SystemHandler) initConfig() {
	if c.initReady {
		return
	}
	c.initReady = true

	c.getConfig()
	engineSystem.Init(c.cb, c.cfg)

	c.srv.ConfigServiceServer = &services.SystemServer{
		EventClient: c.eventClient,
		Cfg:         proto.Clone(c.cfg).(*systempb.Config),
	}
	c.srv.RunServiceServer = &services.SystemServer{
		EventClient: c.eventClient,
	}
}

func (c *SystemHandler) listenEvent() {
	utilsLog.Info("listenEvent : enter")
	for {
		evt, err := c.eventClient.ReceiveEvent()
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		utilsLog.Info("evt = ", evt)
		switch evt.GetType() {
		case eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE:
			c.registerCallBack()
			c.initConfig()
			c.sendReadyEvent()
		case eventpb.InternalTypeOptions_INTERNAL_TYPE_CONFIG:
			if evt.GetConfig().GetActionOption() == eventpb.ConfigActionTypeOptions_CONFIG_ACTION_TYPE_CONFIG_SAVE {
				cfg, _ := c.srv.GetConfig(c.ctx, &emptypb.Empty{})
				c.saveConfig(cfg)
			}
		}
	}
}

func (c *SystemHandler) sendReadyEvent() {
	evt := &eventpb.Internal{
		Ts:      timestamppb.Now(),
		Type:    eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE,
		Message: c.service.String() + eventpb.ServiceActionTypeOptions_SERVICE_ACTION_TYPE_START.String(),
		Parameter: &eventpb.Internal_Init{
			Init: &eventpb.ServiceInitialized{
				ServiceType: c.service,
				Action:      eventpb.ServiceActionTypeOptions_SERVICE_ACTION_TYPE_START,
			},
		},
	}
	c.eventClient.SendEvent(evt)
}

func (c *SystemHandler) GetRegisteredMainService() commonpb.ServicesEnumTypeOptions {
	return commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
}

func (c *SystemHandler) initClient() {
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
	c.eventClient = hal.EventClientInit(c.ctx, c.service, unions, required)
	c.deviceClient = hal.DeviceClientInit(c.ctx, c.service)

	c.configClient = &hal.ConfigClient{
		SavedPath:   fmt.Sprintf("%s/%s.yml", c.deviceClient.GetPath().GetConfigSaved(), c.service.String()),
		DefaultPath: fmt.Sprintf("%s/%s.yml", c.deviceClient.GetPath().GetConfigDefault(), c.service.String()),
	}
}

func (c *SystemHandler) Init(ctx context.Context, grpcSrvConn *grpc.Server) {
	c.ctx = ctx
	c.service = c.GetRegisteredMainService()
	log.Println("Init : c.service = ", c.service)

	c.initClient()

	c.srv = &SystemProtectServiceServer{}
	c.srv.ConfigServiceServer = systempb.UnimplementedConfigServiceServer{}
	c.srv.RunServiceServer = systempb.UnimplementedRunServiceServer{}

	systempb.RegisterConfigServiceServer(grpcSrvConn, c.srv)
	systempb.RegisterRunServiceServer(grpcSrvConn, c.srv)

	go c.listenEvent()
}

func init() {
	registerHandler(new(SystemHandler))
}
