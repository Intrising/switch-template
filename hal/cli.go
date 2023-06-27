package hal

import (
	"context"

	commonpb "github.com/Intrising/intri-type/common"
	clipb "github.com/Intrising/intri-type/core/cli"

	utilsRpc "github.com/Intrising/intri-utils/rpc"
)

type CLIClient struct {
	Ctx            context.Context
	InternalClient clipb.InternalServiceClient
}

func CLIClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *CLIClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CLI)
	return &CLIClient{
		Ctx:            ctx,
		InternalClient: clipb.NewInternalServiceClient(client.GetGrpcClient()),
	}
}

// UpdateAutoLogout:
func (c *CLIClient) UpdateAutoLogout(enabled bool, timeout int32) {
	c.InternalClient.UpdateAutoLogout(c.Ctx, &clipb.AutoLogoutSetting{Enabled: enabled, LogoutTime: timeout})
}
