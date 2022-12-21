package services

import (
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	empty             = &emptypb.Empty{}
	errConfigNotReady = fmt.Errorf("config is not ready")
)
