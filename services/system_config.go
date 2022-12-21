package services

import (
	"context"

	engineSystem "github.com/Intrising/intri-core/engine/system"

	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"

	utilsMisc "github.com/Intrising/intri-utils/misc"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SystemServer) GetConfig(ctx context.Context, in *emptypb.Empty) (*systempb.Config, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg).(*systempb.Config), nil
}

func (s *SystemServer) SetConfig(ctx context.Context, in *systempb.Config) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.cfg == nil {
			return empty, errConfigNotReady
		}
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
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg.GetIdentification()).(*systempb.IdentificationSetting), nil
}

func (s *SystemServer) SetConfigIdentification(ctx context.Context, in *systempb.IdentificationSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.cfg == nil {
			return empty, errConfigNotReady
		}
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	s.SetConfigIdentificationName(ctx, &systempb.ConfigIdentificationName{Value: in.GetName()})
	s.SetConfigIdentificationDescription(ctx, &systempb.ConfigIdentificationDescription{Value: in.GetName()})
	s.SetConfigIdentificationLocation(ctx, &systempb.ConfigIdentificationLocation{Value: in.GetName()})
	s.SetConfigIdentificationContact(ctx, &systempb.ConfigIdentificationContact{Value: in.GetName()})
	if !proto.Equal(s.cfg.GetIdentification(), in) {
		s.cfg.Identification = proto.Clone(in).(*systempb.IdentificationSetting)
		defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
	}

	return empty, nil
}

func (s *SystemServer) GetConfigIdentificationName(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationName, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigIdentificationName{Value: s.cfg.GetIdentification().GetName()}, nil
}

func (s *SystemServer) SetConfigIdentificationName(ctx context.Context, in *systempb.ConfigIdentificationName) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetIdentification().GetName() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Name = in.GetValue()
	_, err := s.SetConfigIdentification(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigIdentificationDescription(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationDescription, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigIdentificationDescription{Value: s.cfg.GetIdentification().GetDescription()}, nil
}

func (s *SystemServer) SetConfigIdentificationDescription(ctx context.Context, in *systempb.ConfigIdentificationDescription) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetIdentification().GetDescription() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Description = in.GetValue()
	_, err := s.SetConfigIdentification(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigIdentificationLocation(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationLocation, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigIdentificationLocation{Value: s.cfg.GetIdentification().GetLocation()}, nil
}

func (s *SystemServer) SetConfigIdentificationLocation(ctx context.Context, in *systempb.ConfigIdentificationLocation) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetIdentification().GetLocation() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Location = in.GetValue()
	_, err := s.SetConfigIdentification(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigIdentificationContact(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigIdentificationContact, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigIdentificationContact{Value: s.cfg.GetIdentification().GetContact()}, nil
}

func (s *SystemServer) SetConfigIdentificationContact(ctx context.Context, in *systempb.ConfigIdentificationContact) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetIdentification().GetContact() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetIdentification()).(*systempb.IdentificationSetting)
	cfg.Contact = in.GetValue()
	_, err := s.SetConfigIdentification(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigOob(ctx context.Context, in *emptypb.Empty) (*systempb.OOBServicePortSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg.GetOob()).(*systempb.OOBServicePortSetting), nil
}

func (s *SystemServer) SetConfigOob(ctx context.Context, in *systempb.OOBServicePortSetting) (*emptypb.Empty, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	return s.SetConfigOobIPv4(ctx, in.GetIPv4())
}

func (s *SystemServer) GetConfigOobIPv4(ctx context.Context, in *emptypb.Empty) (*systempb.OOBServicePortNetworkSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting), nil
}

func (s *SystemServer) SetConfigOobIPv4(ctx context.Context, in *systempb.OOBServicePortNetworkSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.cfg == nil {
			return empty, errConfigNotReady
		}
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	err := engineSystem.SetOOBIPv4(in)
	if err == nil {
		s.SetConfigOobIPv4Enabled(ctx, &systempb.ConfigOobIPv4Enabled{Value: in.GetEnabled()})
		s.SetConfigOobIPv4IPAddr(ctx, &systempb.ConfigOobIPv4IPAddr{Value: in.GetIPAddr()})
		s.SetConfigOobIPv4Netmask(ctx, &systempb.ConfigOobIPv4Netmask{Value: in.GetNetmask()})
		if !proto.Equal(s.cfg.GetOob().GetIPv4(), in) {
			s.cfg.Oob.IPv4 = proto.Clone(in).(*systempb.OOBServicePortNetworkSetting)
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
	}
	return empty, err
}

func (s *SystemServer) GetConfigOobIPv4Enabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4Enabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigOobIPv4Enabled{Value: s.cfg.GetOob().GetIPv4().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigOobIPv4Enabled(ctx context.Context, in *systempb.ConfigOobIPv4Enabled) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetOob().GetIPv4().GetEnabled() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.Enabled = in.GetValue()
	_, err := s.SetConfigOobIPv4(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigOobIPv4IPAddr(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4IPAddr, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigOobIPv4IPAddr{Value: s.cfg.GetOob().GetIPv4().GetIPAddr()}, nil
}

func (s *SystemServer) SetConfigOobIPv4IPAddr(ctx context.Context, in *systempb.ConfigOobIPv4IPAddr) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetOob().GetIPv4().GetIPAddr() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.IPAddr = in.GetValue()
	_, err := s.SetConfigOobIPv4(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigOobIPv4Netmask(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigOobIPv4Netmask, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigOobIPv4Netmask{Value: s.cfg.GetOob().GetIPv4().GetNetmask()}, nil
}

func (s *SystemServer) SetConfigOobIPv4Netmask(ctx context.Context, in *systempb.ConfigOobIPv4Netmask) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetOob().GetIPv4().GetNetmask() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetOob().GetIPv4()).(*systempb.OOBServicePortNetworkSetting)
	cfg.Netmask = in.GetValue()
	_, err := s.SetConfigOobIPv4(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigWatchdog(ctx context.Context, in *emptypb.Empty) (*systempb.WatchDogSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg.GetWatchdog()).(*systempb.WatchDogSetting), nil
}

func (s *SystemServer) SetConfigWatchdog(ctx context.Context, in *systempb.WatchDogSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.cfg == nil {
			return empty, errConfigNotReady
		}
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	err := engineSystem.SetWatchdog(in)
	if err == nil {
		s.SetConfigWatchdogEnabled(ctx, &systempb.ConfigWatchdogEnabled{Value: in.GetEnabled()})
		s.SetConfigWatchdogTriggerTime(ctx, &systempb.ConfigWatchdogTriggerTime{Value: in.GetTriggerTime()})
		if !proto.Equal(s.cfg.GetWatchdog(), in) {
			s.cfg.Watchdog = proto.Clone(in).(*systempb.WatchDogSetting)
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
	}

	return empty, err
}

func (s *SystemServer) GetConfigWatchdogEnabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigWatchdogEnabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigWatchdogEnabled{Value: s.cfg.GetWatchdog().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigWatchdogEnabled(ctx context.Context, in *systempb.ConfigWatchdogEnabled) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetWatchdog().GetEnabled() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetWatchdog()).(*systempb.WatchDogSetting)
	cfg.Enabled = in.GetValue()
	_, err := s.SetConfigWatchdog(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigWatchdogTriggerTime(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigWatchdogTriggerTime, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigWatchdogTriggerTime{Value: s.cfg.GetWatchdog().GetTriggerTime()}, nil
}

func (s *SystemServer) SetConfigWatchdogTriggerTime(ctx context.Context, in *systempb.ConfigWatchdogTriggerTime) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetLogout().GetLogoutTime() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetWatchdog()).(*systempb.WatchDogSetting)
	cfg.TriggerTime = in.GetValue()
	_, err := s.SetConfigWatchdog(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigLogout(ctx context.Context, in *emptypb.Empty) (*systempb.AutoLogoutSetting, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return proto.Clone(s.cfg.GetLogout()).(*systempb.AutoLogoutSetting), nil
}

func (s *SystemServer) SetConfigLogout(ctx context.Context, in *systempb.AutoLogoutSetting) (*emptypb.Empty, error) {
	if !utilsMisc.IsCtxPassLock(ctx) {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.cfg == nil {
			return empty, errConfigNotReady
		}
		ctx = utilsMisc.AddCtxPassLock(ctx)
	}

	s.SetConfigLogoutEnabled(ctx, &systempb.ConfigLogoutEnabled{Value: in.GetEnabled()})
	s.SetConfigLogoutLogoutTime(ctx, &systempb.ConfigLogoutLogoutTime{Value: in.GetLogoutTime()})
	if !proto.Equal(s.cfg.GetLogout(), in) {
		s.cfg.Logout = proto.Clone(in).(*systempb.AutoLogoutSetting)
		defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
	}

	return empty, nil
}

func (s *SystemServer) GetConfigLogoutEnabled(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigLogoutEnabled, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigLogoutEnabled{Value: s.cfg.GetLogout().GetEnabled()}, nil
}

func (s *SystemServer) SetConfigLogoutEnabled(ctx context.Context, in *systempb.ConfigLogoutEnabled) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetLogout().GetEnabled() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetLogout()).(*systempb.AutoLogoutSetting)
	cfg.Enabled = in.GetValue()
	_, err := s.SetConfigLogout(ctx, cfg)
	return empty, err
}

func (s *SystemServer) GetConfigLogoutLogoutTime(ctx context.Context, in *emptypb.Empty) (*systempb.ConfigLogoutLogoutTime, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.cfg == nil {
		return nil, errConfigNotReady
	}
	return &systempb.ConfigLogoutLogoutTime{Value: s.cfg.GetLogout().GetLogoutTime()}, nil
}

func (s *SystemServer) SetConfigLogoutLogoutTime(ctx context.Context, in *systempb.ConfigLogoutLogoutTime) (*emptypb.Empty, error) {
	if utilsMisc.IsCtxPassLock(ctx) {
		if s.cfg.GetLogout().GetLogoutTime() != in.GetValue() {
			defer s.sendConfigChange(in, eventpb.ConfigADUTypeOptions_CONFIG_ADU_TYPE_CONFIG_UPDATE)
		}
		return empty, nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.cfg == nil {
		return empty, errConfigNotReady
	}
	ctx = utilsMisc.AddCtxPassLock(ctx)

	cfg := proto.Clone(s.cfg.GetLogout()).(*systempb.AutoLogoutSetting)
	cfg.LogoutTime = in.GetValue()
	_, err := s.SetConfigLogout(ctx, cfg)
	return empty, err
}
