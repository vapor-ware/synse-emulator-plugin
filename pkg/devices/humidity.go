package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Humidity is the handler for the emulated humidity device(s).
var Humidity = sdk.DeviceHandler{
	Name:  "humidity",
	Read:  humidityRead,
	Write: minMaxCurrentWrite,
}

// humidityRead is the read handler for the emulated humidity device(s).
func humidityRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Humidity.MakeReading(emitter.Next()),
		output.Temperature.MakeReading(emitter.Next()),
	}, nil
}
