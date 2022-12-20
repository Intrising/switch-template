package services

import (
	"sync"
	"time"

	"github.com/Intrising/intri-core/hal"
	systempb "github.com/Intrising/intri-type/core/system"
	eventpb "github.com/Intrising/intri-type/event"
	utilsLog "github.com/Intrising/intri-utils/log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SystemServer struct {
	mutex sync.RWMutex

	systempb.UnimplementedConfigServiceServer
	systempb.UnimplementedInternalServer
	systempb.UnimplementedRunServer

	Cfg         *systempb.Config
	EventClient *hal.EventClient
}

func (s *SystemServer) sendConfigChange(change proto.Message, adu eventpb.ConfigADUTypeOptions) {
	data, err := ptypes.MarshalAny(change)
	if err != nil {
		utilsLog.Error(err)
		return
	}
	evt := &eventpb.Internal{
		Message: "Config Change",
		Type:    eventpb.InternalTypeOptions_INTERNAL_TYPE_CONFIG,
		Ts: &timestamppb.Timestamp{
			Nanos:   int32(time.Now().Nanosecond()),
			Seconds: int64(time.Now().Second()),
		},
		Parameter: &eventpb.Internal_Config{
			Config: &eventpb.ConfigParameter{
				ADUOption:    adu,
				ActionOption: eventpb.ConfigActionTypeOptions_CONFIG_ACTION_TYPE_CONFIG_CHANGE,
				ActionOptionParam: &eventpb.ConfigParameter_ConfigChange{
					ConfigChange: data,
				},
			},
		},
	}

	s.EventClient.SendEvent(evt)
}
