package system

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	// configpb "github.com/Intrising/intri-type/config"
	devicepb "github.com/Intrising/intri-type/device"
	// errorpb "github.com/Intrising/intri-type/error"
	eventpb "github.com/Intrising/intri-type/event"
	// "google.golang.org/protobuf/types/known/emptypb"
)

type CallBack struct {
	Ctx context.Context

	// event
	EventClient interface {
		SendEvent(*eventpb.Internal, string)
	}

	// config
	// ConfigClient interface {
	// 	RunRestoreDefaultConfig(*configpb.RestoreDefaultType) (*emptypb.Empty, error)
	// }

	// device
	DeviceClient interface {
		GetDeviceInfo() *devicepb.Info
		GetPath() *devicepb.PathAll
	}

	// hardware
	HardwareClient interface {
		GetResetButtonStatus() (*commonpb.State, error)
		// SetSystemLED(*hwpb.LEDType) (*emptypb.Empty, error)
	}
}
