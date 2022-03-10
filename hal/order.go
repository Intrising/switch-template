package hal

import (
	"context"
)

type ProtType uint8

const (
	ProtHal ProtType = iota
	ProtHandler
	// YOUR PROT 1
	// YOUR PROT 2
	// YOUR PROT 3
	// YOUR PROT 4
	// ...
	ProtTypeDoNotUseThisEnum
)

// Shaper :
type Shaper interface {
	Init(context.Context)
}

var (
	globalShapesArr map[ProtType]Shaper
)

// RegisterInitCbFunc :
func RegisterInitCbFunc(opt ProtType, initCb Shaper) {
	if globalShapesArr == nil {
		globalShapesArr = make(map[ProtType]Shaper)
	}
	globalShapesArr[opt] = initCb
}

// OrderInit :
func OrderInit() {
	for i := 0; i < int(ProtTypeDoNotUseThisEnum); i++ {
		if ProtType(i) == ProtServices {
			go globalShapesArr[ProtType(i)].Init(halCtx)
		} else {
			globalShapesArr[ProtType(i)].Init(halCtx)
		}
	}
}
