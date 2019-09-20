package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Fan is the handler for the emulated fan device(s).
var Fan = sdk.DeviceHandler{
	Name:  "fan",
	Read:  fanRead,
	Write: fanWrite,
}

// fanRead is the read handler for the emulated fan devices(s).
func fanRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	speed := output.RPM.MakeReading(emitter.Next())
	speed.Context = map[string]string{
		"min": "0",
		"max": "2700",
	}

	return []*output.Reading{
		output.Direction.MakeReading("forward"),
		speed,
	}, nil
}

// fanWrite is the write handler for the emulated fan device(s).
func fanWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	// Parse the data []byte into an int
	v, err := strconv.Atoi(string(data.Data))
	if err != nil {
		return err
	}

	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case "speed":
		emitter.Set(v)
	default:
		return fmt.Errorf("unsupport write action: %v", data.Action)
	}
	return nil
}
