package system

import (
	"fmt"

	systempb "github.com/Intrising/intri-type/core/system"
)

func RegisterCallBack(cb *CallBack, in *systempb.Config) {
	halClient = cb
	cfg = in

	fmt.Println("halClient = ", halClient)
}

func SetAutoLogout(in *systempb.AutoLogoutSetting) error {
	return nil
}

func SetWatchdog(in *systempb.WatchDogSetting) error {
	return nil
}

func SetOOBIPv4(in *systempb.OOBServicePortNetworkSetting) error {
	return nil
}
