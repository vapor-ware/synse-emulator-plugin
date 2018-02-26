package devices

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

// EmulatedHumidity is the handler for the Emulated humidity device.
var EmulatedHumidity = sdk.DeviceHandler{
	Type:  "humidity",
	Model: "emul8-humidity",
	Read:  humidityRead,
}

func humidityRead(device *sdk.Device) ([]*sdk.Reading, error) {
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
