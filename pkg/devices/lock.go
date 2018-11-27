package devices

import (
	"fmt"
	"sync"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	// Valid states of a lock device.
	stateLock   = "locked"
	stateUnlock = "unlocked_electrically"

	// Valid actions of a lock device.
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
		lockState = stateLock
	case actionUnlock:
		lockState = stateUnlock
	case actionPulseUnlock:
		// Unlock the device for 5 seconds then lock it.
		lockState = stateUnlock

		go func() {
			time.Sleep(5 * time.Second)

			mux.Lock()
			defer mux.Unlock()

			lockState = stateLock
		}()
	default:
		return fmt.Errorf("unsupported command for state action: %v", action)
	}

	return nil
}
