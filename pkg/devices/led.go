package devices

import (
	"encoding/hex"
	"fmt"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
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

// ledRead is the read handler for the emulated LED device(s).
func ledRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	val := emitter.Next().(map[string]string)
	state, err := output.State.MakeReading(val["state"])
	if err != nil {
		return nil, err
	}
	color, err := output.Color.MakeReading(val["color"])
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		state,
		color,
	}, nil
}

// ledWrite is the write handler for the emulated LED device(s).
func ledWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	emitter := utils.GetEmitter(device.GetID())
	current := emitter.Next().(map[string]string)

	switch data.Action {
	case "color":
		decoded, err := hex.DecodeString(string(data.Data))
		if err != nil {
			return err
		}
		if len(decoded) != 3 {
			return fmt.Errorf("color value should be a 3-byte (RGB) hex string")
		}
		current["color"] = string(data.Data)
		emitter.Set(current)

	case "state":
		switch cmd := string(data.Data); cmd {
		case stateOn, stateOff, stateBlink:
			current["state"] = cmd
			emitter.Set(current)
		default:
			return fmt.Errorf("unsupported command for state action: %v", cmd)
		}
	default:
		return fmt.Errorf("unsupport write action: %v", data.Action)
	}
	return nil
}
