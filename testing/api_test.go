package test

import (
	"context"
	"log"
	"testing"
	"time"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	systemClient   systempb.ConfigServiceClient
	runClient      systempb.RunClient
	internalClient systempb.InternalClient
	service        = commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
)

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

func init() {
	ctx, _ := context.WithCancel(context.Background())
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM)
	systemClient = systempb.NewConfigServiceClient(client.GetGrpcClient())
	runClient = systempb.NewRunClient(client.GetGrpcClient())
	internalClient = systempb.NewInternalClient(client.GetGrpcClient())
}
