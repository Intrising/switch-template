package test

import (
	"context"
	"log"
	"testing"
	"time"

	commonpb "github.com/Intrising/intri-type/common"
	systempb "github.com/Intrising/intri-type/core/system"
	utilsRpc "github.com/Intrising/intri-utils/rpc"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	systemClient   systempb.ConfigurationServiceClient
	runClient      systempb.RunServicesClient
	internalClient systempb.InternalServicesClient
	service        = commonpb.ServicesEnumTypeOptions_SERVICES_ENUM_TYPE_CORE_SYSTEM
)

func Test_ConfigInit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	cfg := &systempb.Config{
		Identification: &systempb.IdentificationSetting{
			SysName:     "test123",
			SysLocation: "aaaaa",
			SysGroup:    "bbbbb3",
			SysContact:  "ccccc",
		},
	}
	res, err := systemClient.Init(ctx, cfg)
	if err != nil {
		log.Fatalf("error while calling Test_ConfigInit: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from Test_ConfigInit:", res)
	res2, err2 := systemClient.Get_Config(ctx, &emptypb.Empty{})
	if err2 != nil {
		log.Fatalf("error while calling Test_GetConfig: %v \n", err2)
		t.Errorf("got error")
	}
	log.Println("Response from Test_ConfigInit:", res2)
}

func Test_Set_Config_Identification_SysName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	in := &wrappers.StringValue{
		Value: "intrisingok",
	}
	res, err := systemClient.Set_Config_Identification_SysName(ctx, in)
	if err != nil {
		log.Fatalf("error while calling Test_GetConfig: %v \n", err)
		t.Errorf("got error")
	}
	log.Println("Response from Test_Set_Config_Identification_SysName:", res)
	res2, err2 := systemClient.Get_Config(ctx, &emptypb.Empty{})
	if err2 != nil {
		log.Fatalf("error while calling Test_GetConfig: %v \n", err2)
		t.Errorf("got error")
	}
	log.Println("Response from Test_ConfigInit:", res2)
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
