package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Frequency is the handler for the emulated frequency device(s).
var Frequency = sdk.DeviceHandler{
	Name:  "frequency",
	Read:  frequencyRead,
	Write: minMaxCurrentWrite,
}

// frequencyRead is the read handler for the emulated frequency device(s).
func frequencyRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Frequency.MakeReading(emitter.Next()),
	}, nil
}
