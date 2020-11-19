package utils

import (
	"sync"

	"github.com/vapor-ware/synse-sdk/sdk"
)

// This file is used to hold any global references, e.g. to devices or anything else, which can be
// used by the emulator. This is particularly useful for the emulator when it needs to modify values
// of devices which may not be the device explicitly being written to, or if it needs to emulate side
// effects in other devices.

// Global references to Carousel devices. These are needed since the write device is effectively
// a proxy for the read-only device state.
var (
	CarouselGetRackPositionDevice *sdk.Device
	CarouselSetRackPositionDevice *sdk.Device
	CarouselStateMachine          *sdk.Device
	CarouselStatusDevice          *sdk.Device
	CarouselMutex                 sync.Mutex
)
