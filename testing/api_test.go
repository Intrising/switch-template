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
	systemClient   systempb.ConfigurationServiceClient
	runClient      systempb.RunServicesClient
	internalClient systempb.InternalServicesClient
	service        = commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
)

func Test_GetConfig(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res, err := systemClient.Get_Config(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling Test_GetConfig: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from Test_GetConfig:", res)
}

func Test_Run_Get_Info(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res, err := runClient.Run_Get_Info(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling Run_Get_Info: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from Run_Get_Info:", res)
}

func Test_Run_Get_Security_Info(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res, err := internalClient.Run_Get_Security_Info(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling Run_Get_Info: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from Run_Get_Info:", res)
}

func init() {
	ctx, _ := context.WithCancel(context.Background())
	client := utilsRpc.NewClientConn(ctx, service, commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM)
	systemClient = systempb.NewConfigurationServiceClient(client.GetGrpcClient())
	runClient = systempb.NewRunServicesClient(client.GetGrpcClient())
	internalClient = systempb.NewInternalServicesClient(client.GetGrpcClient())
}
