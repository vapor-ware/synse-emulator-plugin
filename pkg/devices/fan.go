package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

var speed int

// Fan is the handler for the emulated fan device(s).
var Fan = sdk.DeviceHandler{
	Name:  "fan",
	Read:  fanRead,
	Write: fanWrite,
}

// fanRead is the read handler for the emulated fan devices(s). It
// returns the `speed` state for the device.
func fanRead(_ *sdk.Device) ([]*output.Reading, error) {
	return []*output.Reading{
		output.RPM.MakeReading(speed),
	}, nil
}

// fanWrite is the write handler for the emulated fan device(s). It
// sets the `speed` state based on the values written to the device.
func fanWrite(_ *sdk.Device, data *sdk.WriteData) error {
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
		speed = s
	}
	return nil
}
