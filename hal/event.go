package hal

import (
	"context"
	"errors"

	commonpb "github.com/Intrising/intri-type/common"
	eventpb "github.com/Intrising/intri-type/event"
	utilsEvent "github.com/Intrising/intri-utils/event"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
)

type EventClient struct {
	Ctx            context.Context
	Client         eventpb.EventClient
	InternalClient *utilsEvent.EventInternal
}

func EventClientInit(ctx context.Context, service commonpb.ServicesEnumTypeOptions) *EventClient {
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_EVENT)
	return &EventClient{
		Ctx: ctx,
		InternalClient: utilsEvent.NewEventInternal(ctx, service,
			[]eventpb.InternalTypeOptions{eventpb.InternalTypeOptions_INTERNAL_TYPE_SYSTEM,
				eventpb.InternalTypeOptions_INTERNAL_TYPE_MAINTENANCE,
				eventpb.InternalTypeOptions_INTERNAL_TYPE_NTP,
				eventpb.InternalTypeOptions_INTERNAL_TYPE_BOOT,
				eventpb.InternalTypeOptions_INTERNAL_TYPE_BUTTON,
				eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE,
			},
			eventpb.ListenTypeOptions_LISTEN_TYPE_BOTH,
		),
		Client: eventpb.NewEventClient(client.GetGrpcClient()),
	}
}

// ReceiveEvent:
func (c *EventClient) ReceiveEvent() (*eventpb.Internal, error) {
	if c.InternalClient == nil {
		return nil, errors.New("event client is nil")
	}
	return c.InternalClient.ReceiveEvent(), nil
}

// SendEvent :
func (c *EventClient) SendEvent(evt *eventpb.Internal, opt string) {
	// if opt == commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_NETWORK.String() {
	// 	if networkEventClient == nil {
	// 		networkEventClient = utilsEvent.NewEventInternal(halCtx, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_NETWORK,
	// 			[]eventpb.InternalTypeOptions{},
	// 			eventpb.ListenTypeOptions_LISTEN_TYPE_TX,
	// 		)
	// 	}
	// 	networkEventClient.SendEvent(evt)
	// } else if opt == commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_TIME.String() {
	// 	if timeEventClient == nil {
	// 		timeEventClient = utilsEvent.NewEventInternal(halCtx, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_TIME,
	// 			[]eventpb.InternalTypeOptions{},
	// 			eventpb.ListenTypeOptions_LISTEN_TYPE_TX,
	// 		)
	// 	}
	// 	timeEventClient.SendEvent(evt)
	// } else {
	if c.InternalClient != nil {
		c.InternalClient.SendEvent(evt)
	}
	// }
}

// EncodeDecode :
func (c *EventClient) EncodeDecode(in *eventpb.CryptoRequest) (*eventpb.CryptoResponse, error) {
	return c.Client.EncodeDecode(c.Ctx, in)
}
