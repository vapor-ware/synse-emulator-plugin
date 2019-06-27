package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Temperature is the handler for the emulated temperature device(s).
var Temperature = sdk.DeviceHandler{
	Name:  "temperature",
	Read:  temperatureRead,
	Write: minMaxCurrentWrite,
}

// temperatureRead is the read handler for the emulated temperature device(s).
func temperatureRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Temperature.MakeReading(emitter.Next()),
	}, nil
}
