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
// been set, this will set the state to 'locked'.
func lockRead(device *sdk.Device) ([]*sdk.Reading, error) {
	mux.Lock()
	defer mux.Unlock()

	var lockState string

	dState, ok := deviceState[device.GUID()]

	if ok {
		if _, ok := dState["lockState"]; ok {
			lockState = dState["lockState"].(string)
		}
	}

	// if we have no stored device lock state, default to "locked"
	if lockState == "" {
		lockState = stateLock
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
func lockWrite(device *sdk.Device, data *sdk.WriteData) error {
	mux.Lock()
	defer mux.Unlock()

	dState, ok := deviceState[device.GUID()]

	switch action := data.Action; action {
	case actionLock:
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockState": stateLock}
		} else {
			dState["lockState"] = stateLock
		}
	case actionUnlock:
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockState": stateUnlock}
		} else {
			dState["lockState"] = stateUnlock
		}
	case actionPulseUnlock:
		// Unlock the device for 5 seconds then lock it.
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"lockState": stateUnlock}
		} else {
			dState["lockState"] = stateUnlock
		}

		go func() {
			time.Sleep(5 * time.Second)

			mux.Lock()
			defer mux.Unlock()

			if !ok {
				deviceState[device.GUID()] = map[string]interface{}{"lockState": stateLock}
			} else {
				dState["lockState"] = stateLock
			}
		}()
	default:
		return fmt.Errorf("unsupported command for state action: %v", action)
	}

	return nil
}
