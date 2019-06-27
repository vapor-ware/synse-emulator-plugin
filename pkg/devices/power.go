package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Power is the handler for the emulated power device(s).
var Power = sdk.DeviceHandler{
	Name:  "power",
	Read:  powerRead,
	Write: minMaxCurrentWrite,
}

// powerRead is the read handler for the emulated power device(s).
func powerRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Watt.MakeReading(emitter.Next()),
	}, nil
}
