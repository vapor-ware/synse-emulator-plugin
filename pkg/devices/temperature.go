package devices

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Temperature is the handler for the emulated temperature device(s).
var Temperature = sdk.DeviceHandler{
	Name:  "temperature",
	Read:  temperatureRead,
	Write: temperatureWrite,
}

// temperatureRead is the read handler for the emulated temperature device(s).
// It returns random values between 0 and 100.
func temperatureRead(device *sdk.Device) ([]*sdk.Reading, error) {
	// Default reading ranges
	var min int = 0
	var max int = 100

	dState, ok := deviceState[device.ID()]
	if ok {
		if _, ok := dState["min"]; ok {
			min = dState["min"].(int)
		}

		if _, ok := dState["max"]; ok {
			max = dState["max"].(int)
		}
	}

	// In the event that we change the min value before the max value to
	// something greater than what the max value was, make the max value
	// bigger so we dont panic when generating a random val.
	if min > max {
		max = min + 1
	}

	logrus.WithFields(logrus.Fields{
		"device": device.ID(),
		"min":    min,
		"max":    max,
	}).Info("reading temp")

	temperature, err := device.GetOutput("temperature").MakeReading(utils.RandIntInRange(min, max))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		temperature,
	}, nil
}

// temperatureWrite is the write handler for the emulated temperature device(s).
// Typically, temperature devices are not writable, but since this is an emulator
// and we may want to change the returned value(s) of a device at runtime, we can
// reset the min and max values.
func temperatureWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == "min" {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		min, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.ID()]
		if !ok {
			deviceState[device.ID()] = map[string]interface{}{"min": min}
		} else {
			dataMap["min"] = min
		}

	} else if action == "max" {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		max, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.ID()]
		if !ok {
			deviceState[device.ID()] = map[string]interface{}{"max": max}
		} else {
			dataMap["max"] = max
		}
	}
	return nil
}
