package devices

import (
	"fmt"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"strconv"

	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Fan is the handler for the emulated fan device(s).
var Fan = sdk.DeviceHandler{
	Name:  "fan",
	Read:  fanRead,
	Write: fanWrite,
}

// fanRead is the read handler for the emulated fan devices(s). It
// returns the `speed` state for the device.
func fanRead(device *sdk.Device) ([]*output.Reading, error) {
	var speed int

	dState, ok := deviceState[device.GetID()]
	if ok {
		if _, ok := dState["speed"]; ok {
			speed = dState["speed"].(int)
		}
	}

	return []*output.Reading{
		outputs.Airflow.MakeReading(speed),
	}, nil
}

// fanWrite is the write handler for the emulated fan device(s). It
// sets the `speed` state based on the values written to the device.
func fanWrite(device *sdk.Device, data *sdk.WriteData) error {
	dState, ok := deviceState[device.GetID()]

	action := data.Action
	raw := data.Data

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == "speed" {
		s, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}

		if !ok {
			deviceState[device.GetID()] = map[string]interface{}{"speed": s}
		} else {
			dState["speed"] = s
		}
	}
	return nil
}
