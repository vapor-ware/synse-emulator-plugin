package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
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
	temp, err := output.Temperature.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		temp,
	}, nil
}
