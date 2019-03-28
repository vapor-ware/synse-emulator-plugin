package devices

import (
	"encoding/hex"
	"fmt"

	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

const (
	stateOn    = "on"
	stateOff   = "off"
	stateBlink = "blink"
)

// LED is the handler for the emulated LED device(s).
var LED = sdk.DeviceHandler{
	Name:  "led",
	Read:  ledRead,
	Write: ledWrite,
}

// ledRead is the read handler for the emulated LED device(s). It
// returns the state and color values for the device.
func ledRead(device *sdk.Device) ([]*output.Reading, error) {
	var state, color string

	dState, ok := deviceState[device.GetID()]

	if ok {
		if _, ok := dState["color"]; ok {
			color = dState["color"].(string)
		}

		if _, ok := dState["state"]; ok {
			state = dState["state"].(string)
		}

	}

	// if we have no stored device led state, default to off
	if state == "" {
		state = stateOff
	}

	// if we have no stored device led color, default to black
	if color == "" {
		color = "000000"
	}

	return []*output.Reading{
		output.State.MakeReading(state),
		output.Color.MakeReading(color),
	}, nil
}

// ledWrite is the write handler for the emulated LED device(s). It
// sets the state and color values for the device.
func ledWrite(device *sdk.Device, data *sdk.WriteData) error { // nolint: gocyclo
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

		dState, ok := deviceState[device.GetID()]
		if !ok {
			deviceState[device.GetID()] = map[string]interface{}{"color": string(raw)}
		} else {
			dState["color"] = string(raw)
		}

	} else if action == "state" {
		cmd := string(raw)
		dState, ok := deviceState[device.GetID()]

		if cmd == stateOn {
			if !ok {
				deviceState[device.GetID()] = map[string]interface{}{"state": stateOn}
			} else {
				dState["state"] = stateOn
			}
		} else if cmd == stateOff {
			if !ok {
				deviceState[device.GetID()] = map[string]interface{}{"state": stateOff}
			} else {
				dState["state"] = stateOff
			}
		} else if cmd == stateBlink {
			if !ok {
				deviceState[device.GetID()] = map[string]interface{}{"state": stateBlink}
			} else {
				dState["state"] = stateBlink
			}
		} else {
			return fmt.Errorf("unsupported command for state action: %v", cmd)
		}
	}
	return nil
}
