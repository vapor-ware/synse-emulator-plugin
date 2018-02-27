package devices

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

// EmulatedAirflow is the handler for the emulated airflow device.
var EmulatedAirflow = sdk.DeviceHandler{
	Type:  "airflow",
	Model: "emul8-air",
	Read:  airflowRead,
}

// airflowRead is the read handler for the emulated airflow device(s). It
// returns random values between the device's min and max range.
func airflowRead(device *sdk.Device) ([]*sdk.Reading, error) {
	var readings []*sdk.Reading
	for _, output := range device.Output {
		min := output.Range.Min
		max := output.Range.Max

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		val := r1.Int31n(max-min) + min

		r := sdk.NewReading(output.Type, strconv.Itoa(int(val)))
		readings = append(readings, r)
	}
	return readings, nil
}
