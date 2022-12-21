package test

import (
	"context"
	"log"
	"testing"
	"time"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"

	utilsEvent "github.com/Intrising/intri-utils/event"
	utilsRpc "github.com/Intrising/intri-utils/rpc"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	systemClient systempb.ConfigServiceClient
	runClient    systempb.RunServiceClient

	eventClient *utilsEvent.EventInternal

	service = commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
)

func Test_SendEventInit(t *testing.T) {
	in := service
	evt := &eventpb.Internal{
		Ts:           timestamppb.Now(),
		Type:         eventpb.InternalTypeOptions_INTERNAL_TYPE_SERVICE,
		ServicesType: in,
		Message:      "The service can start doing initialization",
		LoggingType:  eventpb.LoggingTypeOptions_LOGGING_TYPE_NONE,
		Parameter: &eventpb.Internal_Init{
			Init: &eventpb.ServiceInitialized{
				ServiceType: in,
				Action:      eventpb.ServiceActionTypeOptions_SERVICE_ACTION_TYPE_INIT,
			},
		},
	}
	eventClient.SendEvent(evt)
}

func Test_SetConfigIdentification(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	in := &systempb.IdentificationSetting{
		Name:        "test1",
		Description: "test2",
		Location:    "test3",
		Contact:     "test4",
	}
	res, err := systemClient.SetConfigIdentification(ctx, in)
	if err != nil {
		log.Fatalf("error while calling SetConfigIdentification: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from SetConfigIdentification:", res)
	res2, err2 := systemClient.GetConfigIdentification(ctx, &emptypb.Empty{})
	if err2 != nil {
		log.Fatalf("error while calling GetConfigIdentification: %v \n", err2)
		t.Errorf("got error")
	}
	log.Println("Response from GetConfigIdentification:", res2)
}

func Test_SetConfigIdentificationContact(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	in := &systempb.ConfigIdentificationContact{
		Value: "0912456798",
	}
	res, err := systemClient.SetConfigIdentificationContact(ctx, in)
	if err != nil {
		log.Fatalf("error while calling SetConfigIdentificationContact: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from SetConfigIdentificationContact:", res)
	res2, err2 := systemClient.GetConfigIdentification(ctx, &emptypb.Empty{})
	if err2 != nil {
		log.Fatalf("error while calling GetConfigIdentification: %v \n", err2)
		t.Errorf("got error")
	}
	log.Println("Response from GetConfigIdentification:", res2)
}

func Test_SetConfigLogoutLogoutTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	in := &systempb.ConfigLogoutLogoutTime{
		Value: 60,
	}
	res, err := systemClient.SetConfigLogoutLogoutTime(ctx, in)
	if err != nil {
		log.Fatalf("error while calling SetConfigLogoutLogoutTime: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from SetConfigLogoutLogoutTime:", res)
	res2, err2 := systemClient.GetConfigLogoutLogoutTime(ctx, &emptypb.Empty{})
	if err2 != nil {
		log.Fatalf("error while calling GetConfigLogoutLogoutTime: %v \n", err2)
		t.Errorf("got error")
	}
	log.Println("Response from GetConfigLogoutLogoutTime:", res2)
}

func Test_SendSaveConfigEvent(t *testing.T) {
	log.Println("Test_SendSaveConfigEvent")
	defer log.Println("Test_SendSaveConfigEvent leave")
	evt := &eventpb.Internal{
		Ts:           timestamppb.Now(),
		Type:         eventpb.InternalTypeOptions_INTERNAL_TYPE_CONFIG,
		ServicesType: commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CONFIG,
		Message:      "save config",
		LoggingType:  eventpb.LoggingTypeOptions_LOGGING_TYPE_NONE,
		Parameter: &eventpb.Internal_Config{
			Config: &eventpb.ConfigParameter{
				ActionOption: eventpb.ConfigActionTypeOptions_CONFIG_ACTION_TYPE_CONFIG_SAVE,
			},
		},
	}
	eventClient.SendEvent(evt)

	time.Sleep(time.Second * 10)
}

func init() {
	ctx, _ := context.WithCancel(context.Background())
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM)
	systemClient = systempb.NewConfigServiceClient(client.GetGrpcClient())
	runClient = systempb.NewRunServiceClient(client.GetGrpcClient())
	// internalClient = systempb.NewInternalServiceClient(client.GetGrpcClient())
	eventClient = utilsEvent.NewEventInternal(ctx, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM, []*eventpb.InternalTypeUnionEntry{}, []commonpb.ServicesEnumTypeOptions{})
}
