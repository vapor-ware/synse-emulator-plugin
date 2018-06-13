package devices

import (
	"encoding/hex"
	"fmt"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	stateOn    = "on"
	stateOff   = "off"
	stateBlink = "blink"
)

var (
	state string
	color string
)

// LED is the handler for the emulated LED device(s).
var LED = sdk.DeviceHandler{
	Name:  "led",
	Read:  ledRead,
	Write: ledWrite,
}

// ledRead is the read handler for the emulated LED device(s). It
// returns the state and color values for the device.
func ledRead(device *sdk.Device) ([]*sdk.Reading, error) {
	if state == "" {
		state = stateOff
	}
	if color == "" {
		color = "000000"
	}

	return []*sdk.Reading{
		device.GetOutput("led.state").MakeReading(state),
		device.GetOutput("led.color").MakeReading(color),
	}, nil
}

// ledWrite is the write handler for the emulated LED device(s). It
// sets the state and color values for the device.
func ledWrite(_ *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == "color" {
		decoded, err := hex.DecodeString(string(raw))
		if err != nil {
			return err
		}
		if len(decoded) != 3 {
			return fmt.Errorf("color value should be a 3-byte (RGB) hex string")
		}
		color = string(raw)

	} else if action == "state" {
		cmd := string(raw)
		if cmd == stateOn {
			state = stateOn
		} else if cmd == stateOff {
			state = stateOff
		} else if cmd == stateBlink {
			state = stateBlink
		} else {
			return fmt.Errorf("unsupported command for state action: %v", cmd)
		}
	}
	return nil
}
