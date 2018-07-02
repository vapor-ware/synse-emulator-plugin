package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Airflow is the handler for the emulated airflow device(s).
var Airflow = sdk.DeviceHandler{
	Name: "airflow",
	Read: airflowRead,
}

// airflowRead is the read handler for the emulated airflow device(s). It
// returns random values between -100 and 100.
func airflowRead(device *sdk.Device) ([]*sdk.Reading, error) {
	airflow, err := device.GetOutput("airflow").MakeReading(utils.RandIntInRange(-100, 100))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		airflow,
	}, nil
}
