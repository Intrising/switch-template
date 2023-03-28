package services

import (
	"context"

	system "github.com/Intrising/switch-template/engine/system"

	systempb "github.com/Intrising/intri-type/core/system"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SystemServer) GetInfo(ctx context.Context, in *emptypb.Empty) (*systempb.Info, error) {
	return system.GetInfo(), nil
}
