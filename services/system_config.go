package services

import (
	"context"

	system "github.com/Intrising/switch-template/engine/system"

	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"

	utilsMisc "github.com/Intrising/intri-utils/misc"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SystemServer) GetConfig(ctx context.Context, in *emptypb.Empty) (*systempb.Config, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg).(*systempb.Config), nil
}

func (s *SystemServer) SetConfig(ctx context.Context, in *systempb.Config) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	var err error
	if _, err = s.SetConfigIdentification(ctx, in.GetIdentification()); err != nil {
		return empty, err
	}
	if _, err = s.SetConfigOob(ctx, in.GetOob()); err != nil {
		return empty, err
	}
	if _, err = s.SetConfigWatchdog(ctx, in.GetWatchdog()); err != nil {
		return empty, err
	}
	if _, err = s.SetConfigLogout(ctx, in.GetLogout()); err != nil {
		return empty, err
	}
	return empty, err
}

func (s *SystemServer) GetConfigIdentification(ctx context.Context, in *emptypb.Empty) (*systempb.IdentificationSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg.GetIdentification()).(*systempb.IdentificationSetting), nil
}

func (s *SystemServer) SetConfigIdentification(ctx context.Context, in *systempb.IdentificationSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	if !proto.Equal(s.Cfg.GetIdentification(), in) {
		s.Cfg.Identification = proto.Clone(in).(*systempb.IdentificationSetting)
		defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
	}

	return empty, nil
}

func (s *SystemServer) GetConfigIdentificationName(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationName, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigIdentificationName{Value: s.Cfg.GetIdentification().GetName()}, nil
}

func (s *SystemServer) SetConfigIdentificationName(ctx context.Context, in *systempb.ConfigIdentificationName) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Name = in.GetValue()
	return s.SetConfigIdentification(ctx, cfg)
}

func (s *SystemServer) GetConfigIdentificationDescription(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationDescription, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigIdentificationDescription{Value: s.Cfg.GetIdentification().GetDescription()}, nil
}

func (s *SystemServer) SetConfigIdentificationDescription(ctx context.Context, in *systempb.ConfigIdentificationDescription) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Description = in.GetValue()
	return s.SetConfigIdentification(ctx, cfg)
}

func (s *SystemServer) GetConfigIdentificationLocation(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationLocation, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigIdentificationLocation{Value: s.Cfg.GetIdentification().GetLocation()}, nil
}

func (s *SystemServer) SetConfigIdentificationLocation(ctx context.Context, in *systempb.ConfigIdentificationLocation) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Location = in.GetValue()
	return s.SetConfigIdentification(ctx, cfg)
}

func (s *SystemServer) GetConfigIdentificationContact(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationContact, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigIdentificationContact{Value: s.Cfg.GetIdentification().GetContact()}, nil
}

func (s *SystemServer) SetConfigIdentificationContact(ctx context.Context, in *systempb.ConfigIdentificationContact) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Contact = in.GetValue()
	return s.SetConfigIdentification(ctx, cfg)
}

func (s *SystemServer) GetConfigOob(ctx context.Context, in *emptypb.Empty) (*systempb.OOBServicePortSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg.GetOob()).(*systempb.OOBServicePortSetting), nil
}

func (s *SystemServer) SetConfigOob(ctx context.Context, in *systempb.OOBServicePortSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	return s.SetConfigOobIPv4(ctx, in.GetIPv4())
}

func (s *SystemServer) GetConfigOobIPv4(ctx context.Context, in *emptypb.Empty) (*systempb.OOBServicePortNetworkSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting), nil
}

func (s *SystemServer) SetConfigOobIPv4(ctx context.Context, in *systempb.OOBServicePortNetworkSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	err := system.SetOOBIPv4(in)
	if err == nil {
		if !proto.Equal(s.Cfg.GetOob().GetIPv4(), in) {
			s.Cfg.Oob.IPv4 = proto.Clone(in).(*systempb.OOBServicePortNetworkSetting)
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
	}
	return empty, err
}

func (s *SystemServer) GetConfigOobIPv4Enabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4Enabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigOobIPv4Enabled{Value: s.Cfg.GetOob().GetIPv4().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigOobIPv4Enabled(ctx context.Context, in *systempb.ConfigOobIPv4Enabled) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.Enabled = in.GetValue()
	return s.SetConfigOobIPv4(ctx, cfg)
}

func (s *SystemServer) GetConfigOobIPv4IPAddr(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4IPAddr, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigOobIPv4IPAddr{Value: s.Cfg.GetOob().GetIPv4().GetIPAddr()}, nil
}

func (s *SystemServer) SetConfigOobIPv4IPAddr(ctx context.Context, in *systempb.ConfigOobIPv4IPAddr) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.IPAddr = in.GetValue()
	return s.SetConfigOobIPv4(ctx, cfg)
}

func (s *SystemServer) GetConfigOobIPv4Netmask(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4Netmask, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigOobIPv4Netmask{Value: s.Cfg.GetOob().GetIPv4().GetNetmask()}, nil
}

func (s *SystemServer) SetConfigOobIPv4Netmask(ctx context.Context, in *systempb.ConfigOobIPv4Netmask) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.Netmask = in.GetValue()
	return s.SetConfigOobIPv4(ctx, cfg)
}

func (s *SystemServer) GetConfigOobIPv4DefaultGateway(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4DefaultGateway, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigOobIPv4DefaultGateway{Value: s.Cfg.GetOob().GetIPv4().GetDefaultGateway()}, nil
}

func (s *SystemServer) SetConfigOobIPv4DefaultGateway(ctx context.Context, in *systempb.ConfigOobIPv4DefaultGateway) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.DefaultGateway = in.GetValue()
	return s.SetConfigOobIPv4(ctx, cfg)
}

func (s *SystemServer) GetConfigWatchdog(ctx context.Context, in *emptypb.Empty) (*systempb.WatchDogSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg.GetWatchdog()).(*systempb.WatchDogSetting), nil
}

func (s *SystemServer) SetConfigWatchdog(ctx context.Context, in *systempb.WatchDogSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	err := system.SetWatchdog(in)
	if err == nil {
		if !proto.Equal(s.Cfg.GetWatchdog(), in) {
			s.Cfg.Watchdog = proto.Clone(in).(*systempb.WatchDogSetting)
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
	}

	return empty, err
}

func (s *SystemServer) GetConfigWatchdogEnabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigWatchdogEnabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigWatchdogEnabled{Value: s.Cfg.GetWatchdog().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigWatchdogEnabled(ctx context.Context, in *systempb.ConfigWatchdogEnabled) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetWatchdog()).(*systempb.WatchDogSetting)
	cfg.Enabled = in.GetValue()
	return s.SetConfigWatchdog(ctx, cfg)
}

func (s *SystemServer) GetConfigWatchdogTriggerTime(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigWatchdogTriggerTime, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigWatchdogTriggerTime{Value: s.Cfg.GetWatchdog().GetTriggerTime()}, nil
}

func (s *SystemServer) SetConfigWatchdogTriggerTime(ctx context.Context, in *systempb.ConfigWatchdogTriggerTime) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetWatchdog()).(*systempb.WatchDogSetting)
	cfg.TriggerTime = in.GetValue()
	return s.SetConfigWatchdog(ctx, cfg)
}

func (s *SystemServer) GetConfigLogout(ctx context.Context, in *emptypb.Empty) (*systempb.AutoLogoutSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return proto.Clone(s.Cfg.GetLogout()).(*systempb.AutoLogoutSetting), nil
}

func (s *SystemServer) SetConfigLogout(ctx context.Context, in *systempb.AutoLogoutSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	err := system.SetAutoLogoutTime(in.GetEnabled(), in.GetLogoutTime())
	if err == nil {
		if !proto.Equal(s.Cfg.GetLogout(), in) {
			s.Cfg.Logout = proto.Clone(in).(*systempb.AutoLogoutSetting)
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
	}

	return empty, nil
}

func (s *SystemServer) GetConfigLogoutEnabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigLogoutEnabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigLogoutEnabled{Value: s.Cfg.GetLogout().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigLogoutEnabled(ctx context.Context, in *systempb.ConfigLogoutEnabled) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetLogout()).(*systempb.AutoLogoutSetting)
	cfg.Enabled = in.GetValue()
	return s.SetConfigLogout(ctx, cfg)
}

func (s *SystemServer) GetConfigLogoutLogoutTime(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigLogoutLogoutTime, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return &systempb.ConfigLogoutLogoutTime{Value: s.Cfg.GetLogout().GetLogoutTime()}, nil
}

func (s *SystemServer) SetConfigLogoutLogoutTime(ctx context.Context, in *systempb.ConfigLogoutLogoutTime) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	cfg := proto.Clone(s.Cfg.GetLogout()).(*systempb.AutoLogoutSetting)
	cfg.LogoutTime = in.GetValue()
	return s.SetConfigLogout(ctx, cfg)
}
