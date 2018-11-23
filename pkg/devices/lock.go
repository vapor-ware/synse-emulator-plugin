package devices

import (
	"fmt"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	stateLock        = "lock"
	stateUnlock      = "unlock"
	statePulseUnlock = "pulseUnlock"
)

// lstate holds the lock state.
var lstate string

// Lock is the handler for the emulated Lock device(s).
var Lock = sdk.DeviceHandler{
	Name:  "lock",
	Read:  lockRead,
	Write: lockWrite,
}

// lockRead is the read handler for the emulated Lock device(s). It
// returns the state values for the device.
func lockRead(device *sdk.Device) ([]*sdk.Reading, error) {
	if lstate == "" {
		lstate = stateLock
	}

	stateReading, err := device.GetOutput("lock.state").MakeReading(lstate)
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
	action := data.Action
	raw := data.Data

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == "state" {
		cmd := string(raw)
		if cmd == stateLock {
			lstate = stateLock
		} else if cmd == stateUnlock {
			lstate = stateUnlock
		} else if cmd == statePulseUnlock {
			lstate = statePulseUnlock
		} else {
			return fmt.Errorf("unsupported command for state action: %v", cmd)
		}
	}
	return nil
}
