package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
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
	humidity := emitter.Next()
	h, err := output.Humidity.MakeReading(humidity)
	if err != nil {
		return nil, err
	}
	p, err := output.Percentage.MakeReading(humidity)
	if err != nil {
		return nil, err
	}
	t, err := output.Temperature.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}

	return []*output.Reading{
		h,
		p, // https://vaporio.atlassian.net/browse/VIO-1389
		t,
	}, nil
}
