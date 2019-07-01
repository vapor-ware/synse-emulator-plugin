package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Voltage is the handler for the emulated voltage device(s).
var Voltage = sdk.DeviceHandler{
	Name:  "voltage",
	Read:  voltageRead,
	Write: minMaxCurrentWrite,
}

// voltageRead is the read handler for the emulated voltage device(s).
func voltageRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Voltage.MakeReading(emitter.Next()),
	}, nil
}
