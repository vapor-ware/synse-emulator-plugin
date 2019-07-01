package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Energy is the handler for the emulated energy device(s).
var Energy = sdk.DeviceHandler{
	Name:  "energy",
	Read:  energyRead,
	Write: minMaxCurrentWrite,
}

// energyRead is the read handler for the emulated energy device(s).
func energyRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.KilowattHour.MakeReading(emitter.Next()),
	}, nil
}
