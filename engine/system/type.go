package system

import (
	"context"

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
		// GetResetButtonStatus() (*commonpb.BoolValue, error)
		// SetSystemLED(*hwpb.LEDType) (*emptypb.Empty, error)
	}
}
