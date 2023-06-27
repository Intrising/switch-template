package hal

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	timepb "github.com/Intrising/intri-type/core/time"

	utilsRpc "github.com/Intrising/intri-utils/rpc"
)

type TimeClient struct {
	Ctx            context.Context
	RunClient      timepb.RunServiceClient
	InternalClient timepb.InternalServiceClient
}

func TimeClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *TimeClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_TIME)
	return &TimeClient{
		Ctx:            ctx,
		InternalClient: timepb.NewInternalServiceClient(client.GetGrpcClient()),
		RunClient:      timepb.NewRunServiceClient(client.GetGrpcClient()),
	}
}

// GetTimeDate:
func (c *TimeClient) GetCorrectionTimeWithString(in string) string {
	val, _ := c.InternalClient.GetCorrectionTimeWithString(c.Ctx, &timepb.RequestWithString{Ts: in})
	return val.GetTs()
}

// GetTimeDate:
func (c *TimeClient) GetTimeDate() *timepb.DateTime {
	val, _ := c.RunClient.GetTimeDate(c.Ctx, empty)
	return val
}
