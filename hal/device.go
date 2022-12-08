package hal

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	devicepb "github.com/Intrising/intri-type/device"

	utilsRpc "github.com/Intrising/intri-utils/rpc"
)

type DeviceClient struct {
	Ctx    context.Context
	Client devicepb.DeviceClient
	Value  struct {
		deviceInfo *devicepb.Info
		pathAll    *devicepb.PathAll
		boundary   *devicepb.BoundaryAll
	}
}

func DeviceClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *DeviceClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_DEVICE)
	return &DeviceClient{
		Ctx:    ctx,
		Client: devicepb.NewDeviceClient(client.GetGrpcClient()),
	}
}

func (c *DeviceClient) getBoundary() {
	for {
		if val, err := c.Client.GetBoundary(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.boundary = val
			break
		}
	}
}

func (c *DeviceClient) getPath() {
	for {
		if val, err := c.Client.GetPath(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.pathAll = val
			break
		}
	}
}

func (c *DeviceClient) getDeviceInfo() {
	for {
		if val, err := c.Client.GetDeviceInfo(c.Ctx, empty); err != nil {
			continue
		} else {
			c.Value.deviceInfo = val
			break
		}
	}
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
