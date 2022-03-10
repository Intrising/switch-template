package hal

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	myService    = "SERVICE_SWIX_TEMPLATE"
	shortTimeout = time.Duration(3) * time.Second
	midTimeout   = time.Duration(10) * time.Second
	longTimeout  = time.Duration(2) * time.Minute
)

var (
	empty     = &emptypb.Empty{}
	halCtx    context.Context
	halCancel context.CancelFunc
)

// HalInit :
type HalInit struct{}

// Init :
func (h HalInit) Init(parent context.Context) {
	fmt.Println("hal Init")

	// start here
}

func init() {
	halCtx, halCancel = context.WithCancel(context.Background())
	RegisterInitCbFunc(ProtHal, new(HalInit))
}
