package hal

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	devicepb "github.com/Intrising/intri-type/device"

	utilsRpc "github.com/Intrising/intri-utils/rpc"
)

type DeviceClient struct {
	Ctx    context.Context
	Client devicepb.RunServiceClient
	Value  struct {
		deviceInfo *devicepb.Info
		pathAll    *devicepb.PathAll
		boundary   *devicepb.BoundaryAll
		boradInfo  *devicepb.BoardInfo
		fcl        *devicepb.FunctionControlAll
	}
}

func DeviceClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *DeviceClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_DEVICE)
	return &DeviceClient{
		Ctx:    ctx,
		Client: devicepb.NewRunServiceClient(client.GetGrpcClient()),
	}
}

func (c *DeviceClient) getFcl() {
	for {
		if val, err := c.Client.GetInfoFunctionControl(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.fcl = val
			break
		}
	}
}

func (c *DeviceClient) getBoundary() {
	for {
		if val, err := c.Client.GetInfoBoundary(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.boundary = val
			break
		}
	}
}

func (c *DeviceClient) getBoardInfo() {
	for {
		if val, err := c.Client.GetInfoBoard(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.boradInfo = val
			break
		}
	}
}

func (c *DeviceClient) getPath() {
	for {
		if val, err := c.Client.GetInfoPath(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.pathAll = val
			break
		}
	}
}

func (c *DeviceClient) getDeviceInfo() {
	for {
		if val, err := c.Client.GetInfoDevice(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.deviceInfo = val
			break
		}
	}
}

// GetBoundary :
func (c *DeviceClient) GetInfoFunctionControl() *devicepb.FunctionControlAll {
	if c.Value.boundary == nil {
		c.getFcl()
	}

	return c.Value.fcl
}

// GetBoundary :
func (c *DeviceClient) GetBoundary() *devicepb.BoundaryAll {
	if c.Value.boundary == nil {
		c.getBoundary()
	}

	return c.Value.boundary
}

// GetDeviceInfo :
func (c *DeviceClient) GetDeviceInfo() *devicepb.Info {
	if c.Value.deviceInfo == nil {
		c.getDeviceInfo()
	}
	return c.Value.deviceInfo
}

// GetPath :
func (c *DeviceClient) GetPath() *devicepb.PathAll {
	if c.Value.pathAll == nil {
		c.getPath()
	}
	return c.Value.pathAll
}

// GetPath :
func (c *DeviceClient) GetBoardInfo() *devicepb.BoardInfo {
	if c.Value.boradInfo == nil {
		c.getBoardInfo()
	}
	return c.Value.boradInfo
}
