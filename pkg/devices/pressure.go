package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Pressure is the handler for the emulated pressure device(s).
var Pressure = sdk.DeviceHandler{
	Name:  "pressure",
	Read:  pressureRead,
	Write: minMaxCurrentWrite,
}

// pressureRead is the read handler for the emulated pressure device(s).
func pressureRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Pascal.MakeReading(emitter.Next()),
	}, nil
}
