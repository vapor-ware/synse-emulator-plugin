package devices

import (
	"fmt"
	"time"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

const (
	// Valid statuses of a lock device.
	statusLock   = "locked"
	statusUnlock = "unlocked_electrically"

	// Valid actions of a lock device.
	actionLock        = "lock"
	actionUnlock      = "unlock"
	actionPulseUnlock = "pulseUnlock"
)

// Lock is the handler for the emulated lock device(s).
var Lock = sdk.DeviceHandler{
	Name:  "lock",
	Read:  lockRead,
	Write: lockWrite,
}

// lockRead is the read handler for the emulated lock device(s).
func lockRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	status, err := output.Status.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		status,
	}, nil
}

// lockWrite is the write handler for the emulated lock device(s).
func lockWrite(device *sdk.Device, data *sdk.WriteData) error {
	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case actionLock:
		emitter.Set(statusLock)
	case actionUnlock:
		emitter.Set(statusUnlock)
	case actionPulseUnlock:
		emitter.Set(statusUnlock)
		go func() {
			time.Sleep(5 * time.Second)
			emitter.Set(statusLock)
		}()
	default:
		return fmt.Errorf("unsupported action for lock write: %v", data.Action)
	}
	return nil
}
