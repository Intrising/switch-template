package system

import (
	"fmt"


	systempb "github.com/Intrising/intri-type/core/system"
)

func Init(cb *CallBack, in *systempb.Config) {
	halClient = cb
	cfg = in
}

func GetInfo() *systempb.Info {
	return &systempb.Info{
	}
}

func SetWatchdog(in *systempb.WatchDogSetting) error {
	if !halClient.DeviceClient.GetInfoFunctionControl().GetSystem().GetWatchdog() {
		return fmt.Errorf("the feature is not supported")
	}
	return nil
}

func SetOOBIPv4(in *systempb.OOBServicePortNetworkSetting) error {
	if !halClient.DeviceClient.GetInfoFunctionControl().GetSystem().GetOob() {
		return fmt.Errorf("the feature is not supported")
	}

	return nil
}
