package system

import (
	"os"

	systempb "github.com/Intrising/intri-type/core/system"

	"golang.org/x/sys/unix"
)

var (
	halClient *CallBack
	cfg       *systempb.Config
)

func getSysUpTime() int32 {
	var si unix.Sysinfo_t
	unix.Sysinfo(&si)
	return (int32)(si.Uptime)
}

func touchBootReady() {
	fileName := halClient.DeviceClient.GetPath().GetBootReady()
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, _ := os.Create(fileName)
		defer file.Close()
	}
}
