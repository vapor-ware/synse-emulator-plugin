package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// KillaWattHours is the handler for the emulated kwh device(s).
var KillaWattHours = sdk.DeviceHandler{
	Name:  "killawatthour",
	Read:  kwhRead,
	Write: kwhWrite,
}

// Power is the handler for the emulated power device(s).
var Power = sdk.DeviceHandler{
	Name:  "power",
	Read:  powerRead,
	Write: powerWrite,
}

// Voltage is the handler for the emulated voltage device(s).
var Voltage = sdk.DeviceHandler{
	Name:  "voltage",
	Read:  voltageRead,
	Write: voltageWrite,
}

// kwhRead is the read handler for the emulated KillaWatt device(s).
// It returns random values between 0 and 100.
func kwhRead(device *sdk.Device) ([]*sdk.Reading, error) {
	// Default reading ranges
	var min = 0
	var max = 4500

	dState, ok := deviceState[device.GUID()]
	if ok {
		if _, ok := dState[MIN]; ok {
			min = dState[MIN].(int)
		}

		if _, ok := dState[MAX]; ok {
			max = dState[MAX].(int)
		}
	}

	// In the event that we change the min value before the max value to
	// something greater than what the max value was, make the max value
	// bigger so we dont panic when generating a random val.
	if min > max {
		max = min + 1
	}

	power, err := device.GetOutput("kwh.power").MakeReading(utils.RandIntInRange(min, max))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		power,
	}, nil
}

// kwhWrite is the write handler for the emulated killawatt device(s).
// Typically, killawatt devices are not writable, but since this is an emulator
// and we may want to change the returned value(s) of a device at runtime, we can
// reset the min and max values.
func kwhWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == MIN {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		min, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MIN: min}
		} else {
			dataMap[MIN] = min
		}

	} else if action == MAX {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		max, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MAX: max}
		} else {
			dataMap[MAX] = max
		}
	}
	return nil
}

// powerRead is the read handler for the emulated power device(s).
// It returns random values between 1000 and 3000.
func powerRead(device *sdk.Device) ([]*sdk.Reading, error) {
	// Default reading ranges
	var min = 1000
	var max = 3000

	dState, ok := deviceState[device.GUID()]
	if ok {
		if _, ok := dState[MIN]; ok {
			min = dState[MIN].(int)
		}

		if _, ok := dState[MAX]; ok {
			max = dState[MAX].(int)
		}
	}

	// In the event that we change the min value before the max value to
	// something greater than what the max value was, make the max value
	// bigger so we dont panic when generating a random val.
	if min > max {
		max = min + 1
	}

	power, err := device.GetOutput("power").MakeReading(utils.RandIntInRange(min, max))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		power,
	}, nil
}

// powerWrite is the write handler for the emulated power device(s).
// Typically, power devices are not writable, but since this is an emulator
// and we may want to change the returned value(s) of a device at runtime, we can
// reset the min and max values.
func powerWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == MIN {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		min, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MIN: min}
		} else {
			dataMap[MIN] = min
		}

	} else if action == MAX {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		max, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MAX: max}
		} else {
			dataMap[MAX] = max
		}
	}
	return nil
}

// voltageRead is the read handler for the emulated power device(s).
// It returns random values between 100 and 500
func voltageRead(device *sdk.Device) ([]*sdk.Reading, error) {
	// Default reading ranges
	var min = 100
	var max = 500

	dState, ok := deviceState[device.GUID()]
	if ok {
		if _, ok := dState[MIN]; ok {
			min = dState[MIN].(int)
		}

		if _, ok := dState[MAX]; ok {
			max = dState[MAX].(int)
		}
	}

	// In the event that we change the min value before the max value to
	// something greater than what the max value was, make the max value
	// bigger so we dont panic when generating a random val.
	if min > max {
		max = min + 1
	}

	power, err := device.GetOutput("voltage").MakeReading(utils.RandIntInRange(min, max))
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		power,
	}, nil
}

// voltageWrite is the write handler for the emulated voltage device(s).
// Typically, voltage devices are not writable, but since this is an emulator
// and we may want to change the returned value(s) of a device at runtime, we can
// reset the min and max values.
func voltageWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == MIN {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		min, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MIN: min}
		} else {
			dataMap[MIN] = min
		}

	} else if action == MAX {
		// This could get dicey, but since `raw` is bytes and synse server basically just
		// encodes it as a string, the int value we expect here is actually the bytes for
		// the string representation of the int...
		max, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}
		dataMap, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{MAX: max}
		} else {
			dataMap[MAX] = max
		}
	}
	return nil
}
