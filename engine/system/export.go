package system

import (
	"fmt"
	"time"

	utilsMisc "github.com/Intrising/intri-utils/misc"
	"google.golang.org/protobuf/proto"

	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"
)

var (
	bootTime time.Time
)

func Init(cb *CallBack, in *systempb.Config) {
	halClient = cb
	cfg = in

	_ = SetWatchdog(in.GetWatchdog())
	_ = SetOOBIPv4(in.GetOob().GetIPv4())
}

func GetInfo() *systempb.Info {
	info := halClient.DeviceClient.GetDeviceInfo()
	boardInfo := halClient.DeviceClient.GetBoardInfo()

	date := halClient.TimeClient.GetTimeDate()
	timeDate := time.Date(int(date.GetYear()), time.Month(date.GetMonth()), int(date.GetDay()), int(date.GetHours()), int(date.GetMinutes()), int(date.GetSeconds()), 0, time.UTC)
	return &systempb.Info{
		Oid:                 info.GetOid(),
		MACAddress:          info.GetMACAddr(),
		SoftwareVersion:     info.GetCurrentSwVersion(),
		HardwareModel:       info.GetModel(),
		HardwareDescription: boardInfo.GetSystemDescription(),
		SerialNumber:        info.GetSerialNo(),
		UpTime:              getSysUpTime(),
		LayerMode:           info.GetLayer(),
		Date:                timeDate.Format("Mon Jan _2 15:04:05 MST 2006"), // wait time
	}
}

func SetWatchdog(in *systempb.WatchDogSetting) error {
	if !halClient.DeviceClient.GetInfoFunctionControl().GetSystem().GetWatchdog() {
		return fmt.Errorf("the feature is not supported")
	}
	if cfg.Watchdog.Enabled != in.GetEnabled() {
		cfg.Watchdog.Enabled = in.GetEnabled()
		if in.GetEnabled() {
			utilsMisc.ShellExec("/usr/sbin/watchdog.sh start %d", in.GetTriggerTime())
		} else {
			utilsMisc.ShellExec("/usr/sbin/watchdog.sh stop")
		}
	}
	return nil
}

func SetOOBIPv4(in *systempb.OOBServicePortNetworkSetting) error {
	if !halClient.DeviceClient.GetInfoFunctionControl().GetSystem().GetOob() {
		return fmt.Errorf("the feature is not supported")
	}

	if !proto.Equal(cfg.GetOob(), in) {
		cfg.Oob = proto.Clone(in).(*systempb.OOBServicePortSetting)

		utilsMisc.ShellExec("/sbin/set_service_port_ip.sh delv4")
		utilsMisc.ShellExec("/sbin/set_service_port_ip.sh cleargw")
		if !in.GetEnabled() || cfg.GetOob().GetIPv4().GetIPAddr() == "" || cfg.GetOob().GetIPv4().GetNetmask() == "" {
			return nil
		}

		time.Sleep(time.Second * time.Duration(1))

		cidr, err := ConvertIPMaskToCIDR(cfg.GetOob().GetIPv4().GetIPAddr(), cfg.GetOob().GetIPv4().GetNetmask())
		if err != nil {
			return err
		}
		utilsMisc.ShellExec("/sbin/set_service_port_ip.sh addv4 %s %s", cfg.GetOob().GetIPv4().GetIPAddr(), cfg.GetOob().GetIPv4().GetNetmask())
		utilsMisc.ShellExec("/sbin/set_service_port_ip.sh addgw %s %s %s", cfg.GetOob().GetIPv4().GetIPAddr(), cidr, cfg.GetOob().GetIPv4().GetDefaultGateway())
	}

	return nil
}

func GetOOBIPv4Status() (*systempb.OOBServiceStatus, error) {
	if !halClient.DeviceClient.GetInfoFunctionControl().GetSystem().GetOob() {
		return nil, fmt.Errorf("the feature is not supported")
	}

	if !cfg.GetOob().GetIPv4().GetEnabled() {
		return nil, fmt.Errorf("OOB Service Port is currently disabled")
	}

	var oobServiceMACAddr string
	var err error
	if oobServiceMACAddr, err = utilsMisc.ShellExec("cat %s | tr -d '\n'", halClient.DeviceClient.GetPath().GetOOBServicePortMAC()); err != nil {
		oobServiceMACAddr = ""
	}
	return &systempb.OOBServiceStatus{
		IPAddr:         cfg.GetOob().GetIPv4().GetIPAddr(),
		Netmask:        cfg.GetOob().GetIPv4().GetNetmask(),
		DefaultGateway: cfg.GetOob().GetIPv4().GetDefaultGateway(),
		MACAddress:     oobServiceMACAddr,
	}, nil
}

// ExecuteSystemReadyEvent :
func ExecuteSystemReadyEvent() {
	touchBootReady()
	sendBootEvent(eventpb.BootActionTypeOptions_BOOT_ACTION_TYPE_READY, "")
}

// ExecuteBootEvent :
func ExecuteBootEvent() {
	fmt.Println("ExecuteBootEvent")
	warmStartPath := halClient.DeviceClient.GetPath().GetWarmStart()
	version := halClient.DeviceClient.GetDeviceInfo().GetCurrentSwVersion()
	if utilsMisc.CheckFileExists(warmStartPath) {
		utilsMisc.RemoveFile(warmStartPath)
		sendBootEvent(eventpb.BootActionTypeOptions_BOOT_ACTION_TYPE_WARM_START, version)
	} else {
		sendBootEvent(eventpb.BootActionTypeOptions_BOOT_ACTION_TYPE_COLD_START, version)
	}

	bootTime = time.Now()
	fmt.Println("bootTime = ", bootTime)
}

func HandlerResetButton(in *eventpb.ButtonParameter) {
}

func SetAutoLogoutTime(enabled bool, timeout int32) error {
	if halClient.CLIClient == nil {
		return fmt.Errorf("the cli client is not ready")
	}
	halClient.CLIClient.UpdateAutoLogout(enabled, timeout)
	return nil
}
