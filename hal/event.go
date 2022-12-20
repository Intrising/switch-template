package hal

import (
	"context"
	"errors"
	"fmt"

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
	fmt.Println("EventClientInit : enter")
	defer fmt.Println("EventClientInit : leave")
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_EVENT)

	unions := []*eventpb.InternalTypeUnionEntry{
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_SYSTEM},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_NTP},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_BOOT},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_BUTTON},
		{Option: eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE},
	}

	required := []commonpb.ServicesEnumTypeOptions{}

	return &EventClient{
		Ctx:            ctx,
		InternalClient: utilsEvent.NewEventInternal(ctx, service, unions, required),
		Client:         eventpb.NewEventClient(client.GetGrpcClient()),
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
func (c *EventClient) SendEvent(evt *eventpb.Internal) {
	fmt.Println("SendEvent : ", evt)
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
	fmt.Println("c = ", c)
	fmt.Println("Client = ", c.Client)
	fmt.Println("c.InternalClient = ", c.InternalClient == nil)

	if c.InternalClient != nil {
		c.InternalClient.SendEvent(evt)
	}
	// }
}

// EncodeDecode :
func (c *EventClient) EncodeDecode(in *eventpb.CryptoRequest) (*eventpb.CryptoResponse, error) {
	return c.Client.EncodeDecode(c.Ctx, in)
}
