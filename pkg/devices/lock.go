package devices

import (
	"fmt"
	"sync"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	actionLock        = "lock"
	actionUnlock      = "unlock"
	actionPulseUnlock = "pulseUnlock"
)

var (
	// lockState holds the lock state.
	lockState string

	// mux provides mutual exclusion for reading/writing to lock state.
	mux sync.Mutex
)

// Lock is the handler for the emulated Lock device(s).
var Lock = sdk.DeviceHandler{
	Name:  "lock",
	Read:  lockRead,
	Write: lockWrite,
}

// lockRead is the read handler for the emulated Lock device(s). It
// returns the state values for the device. If no state has previously
// been set, this will set the state to 'lock'
func lockRead(device *sdk.Device) ([]*sdk.Reading, error) {
	mux.Lock()
	defer mux.Unlock()

	if lockState == "" {
		lockState = actionLock
	}

	stateReading, err := device.GetOutput("lock.state").MakeReading(lockState)
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		stateReading,
	}, nil
}

// lockWrite is the write handler for the emulated Lock device(s). It
// sets the state values for the device.
func lockWrite(_ *sdk.Device, data *sdk.WriteData) error {
	mux.Lock()
	defer mux.Unlock()

	switch action := data.Action; action {
	case actionLock:
		lockState = "lock"
	case actionUnlock:
		lockState = "unlock"
	case actionPulseUnlock:
		// Unlock the device for 5 seconds then lock it.
		lockState = "unlock"

		go func() {
			time.Sleep(5 * time.Second)

			mux.Lock()
			defer mux.Unlock()

			lockState = "lock"
		}()
	default:
		return fmt.Errorf("unsupported command for state action: %v", action)
	}

	return nil
}
