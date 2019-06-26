package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Pressure is the handler for the emulated pressure device(s).
var Pressure = sdk.DeviceHandler{
	Name:  "pressure",
	Read:  pressureRead,
	Write: pressureWrite,
}

// pressureRead is the read handler for the emulated pressure device(s). It
// returns random values between -5 and 5
func pressureRead(device *sdk.Device) ([]*output.Reading, error) {
	// Default reading ranges
	var min = -5
	var max = 5

	dState, ok := deviceState[device.GetID()]
	if ok {
		if _, ok := dState[MIN]; ok {
			min = dState[MIN].(int)
		}

		if _, ok := dState[MAX]; ok {
			max = dState[MAX].(int)
		}
	}

	// In the event that we change the min value before the max value to
	// something greater than what the max value was, make the max value
	// bigger so we dont panic when generating a random val.
	if min > max {
		max = min + 1
	}

	return []*output.Reading{
		output.Pascal.MakeReading(utils.RandIntInRange(min, max)),
	}, nil
}

// pressureWrite is the write handler for the emulated pressure device(s).
// Typically, pressure devices are not writable, but since this is an emulator
// and we may want to change the returned value(s) of a device at runtime, we can
// reset the min and max values.
func pressureWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == MIN {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		min, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GetID()]
		if !ok {
			deviceState[device.GetID()] = map[string]interface{}{MIN: min}
		} else {
			dataMap[MIN] = min
		}

	} else if action == MAX {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		max, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GetID()]
		if !ok {
			deviceState[device.GetID()] = map[string]interface{}{MAX: max}
		} else {
			dataMap[MAX] = max
		}
	}
	return nil
}
