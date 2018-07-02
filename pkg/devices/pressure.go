package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Pressure is the handler for the emulated pressure device(s).
var Pressure = sdk.DeviceHandler{
	Name: "pressure",
	Read: pressureRead,
}

// pressureRead is the read handler for the emulated pressure device(s). It
// returns random values between -5 and 5
func pressureRead(device *sdk.Device) ([]*sdk.Reading, error) {
	pressure, err := device.GetOutput("pressure").MakeReading(utils.RandIntInRange(-5, 5))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		pressure,
	}, nil
}
