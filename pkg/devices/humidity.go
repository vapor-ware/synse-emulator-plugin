package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Humidity is the handler for the emulated humidity device(s).
var Humidity = sdk.DeviceHandler{
	Name: "humidity",
	Read: humidityRead,
}

// humidityRead is the read handler for the emulated humidity device(s). It
// returns random values between 0 and 100.
func humidityRead(device *sdk.Device) ([]*sdk.Reading, error) {
	return []*sdk.Reading{
		device.GetOutput("humidity").MakeReading(utils.RandIntInRange(0, 100)),
		device.GetOutput("temperature").MakeReading(utils.RandIntInRange(0, 100)),
	}, nil
}
