package handler

import (
	"context"

	"github.com/Intrising/switch-template/hal"
	_ "github.com/Intrising/switch-template/services"
)

func registerCallbackFunction() {

}

type HandlerInit struct{}

// Init :
func (p HandlerInit) Init(ctx context.Context) {
	registerCallbackFunction()
	// go listenReceiveEvent()
}

func init() {
	hal.RegisterInitCbFunc(hal.ProtHandler, new(HandlerInit))
}
