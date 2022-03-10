package services

import (
	"context"
	"fmt"

	"github.com/Intrising/switch-template/hal"

	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	empty       = &emptypb.Empty{}
	servicesCtx context.Context
)

// ServicesInit :
type ServicesInit struct{}

// Init:
func (s ServicesInit) Init(parent context.Context) {
	fmt.Println("services Init")
	servicesCtx = parent

	// start here
}

func init() {
	hal.RegisterInitCbFunc(hal.ProtServices, new(ServicesInit))
}
