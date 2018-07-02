package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Temperature is the handler for the emulated temperature device(s).
var Temperature = sdk.DeviceHandler{
	Name: "temperature",
	Read: temperatureRead,
}

// temperatureRead is the read handler for the emulated temperature device(s).
// It returns random values between 0 and 100.
func temperatureRead(device *sdk.Device) ([]*sdk.Reading, error) {
	temperature, err := device.GetOutput("temperature").MakeReading(utils.RandIntInRange(0, 100))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		temperature,
	}, nil
}
