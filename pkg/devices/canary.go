package devices

import (
	"fmt"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

// power is the handler for the emulated power device(s).
var Canary = sdk.DeviceHandler{
	Name:  "power",
	Read:  canaryRead,
	Write: canaryWrite,
}

// powerRead is the read handler for the emulated power device(s).
func canaryRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	val := emitter.Next().(map[string]string)
	state, err := output.State.MakeReading(val["state"])
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		state,
	}, nil
}

// powerWrite is the write handler for the emulated LED device(s).
func canaryWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	emitter := utils.GetEmitter(device.GetID())
	current := emitter.Next().(map[string]string)

	switch data.Action {
	case "state":
		current["state"] = string(data.Data)
		emitter.Set(current)
	default:
		return fmt.Errorf("unsupport write action: %v", data.Action)
	}
	return nil
}
