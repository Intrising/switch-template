package system

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	configpb "github.com/Intrising/intri-type/config"
	devicepb "github.com/Intrising/intri-type/device"
	eventpb "github.com/Intrising/intri-type/event"

	timepb "github.com/Intrising/intri-type/core/time"
)

type CallBack struct {
	Ctx context.Context

	// event
	EventClient interface {
		SendEvent(*eventpb.Internal)
		ReceiveEvent() (*eventpb.Internal, error)
	}

	// device
	DeviceClient interface {
		GetDeviceInfo() *devicepb.Info
		GetPath() *devicepb.PathAll
		GetBoardInfo() *devicepb.BoardInfo
		GetInfoFunctionControl() *devicepb.FunctionControlAll
	}

	// hardware
	HardwareClient interface {
		GetResetButtonStatus() (*commonpb.BoolValue, error)
		// SetSystemLED(*hwpb.LEDType) (*emptypb.Empty, error)
	}

	ConfigClient interface {
		RestoreDefault(in configpb.FactoryDefaultModeTypeOptions) error
	}

	TimeClient interface {
		GetTimeDate() *timepb.DateTime
	}

	CLIClient interface {
		UpdateAutoLogout(bool, int32)
	}
}
