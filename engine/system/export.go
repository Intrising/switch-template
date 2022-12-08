package system

import (
	systempb "github.com/Intrising/intri-type/core/system"
)

func RegisterCallBack(in *CallBack) {
	halClient = in
}

func SetLogout(in *systempb.AutoLogoutSetting) {

}
