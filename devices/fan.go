package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-sdk/sdk"
)

var speed int

// EmulatedFan is the handler for the emulated fan device.
var EmulatedFan = sdk.DeviceHandler{
	Type:  "fan",
	Model: "emul8-fan",
	Read:  fanRead,
	Write: fanWrite,
}

// fanRead is the read handler for the emulated fan devices(s). It
// returns the `speed` state for the device.
func fanRead(device *sdk.Device) ([]*sdk.Reading, error) {
	ret := []*sdk.Reading{
		sdk.NewReading("fan_speed", strconv.Itoa(speed)),
	}
	return ret, nil
}

// airflowWrite is the write handler for the emulated fan device(s). It
// sets the `speed` state based on the values written to the device.
func fanWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "speed" {
		s, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			return err
		}
		speed = s
	}
	return nil
}
