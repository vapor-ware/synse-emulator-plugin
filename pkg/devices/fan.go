package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Fan is the handler for the emulated fan device(s).
var Fan = sdk.DeviceHandler{
	Name:  "fan",
	Read:  fanRead,
	Write: fanWrite,
}

// FanMulti is the handler for emulated fan device(s) which provide more
// than one RPM reading value.
var FanMulti = sdk.DeviceHandler{
	Name:  "fan-multi",
	Read:  fanMultiRead,
	Write: fanMultiWrite,
}

// fanRead is the read handler for the emulated fan devices(s).
func fanRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	speed := output.RPM.MakeReading(emitter.Next())
	speed.Context = map[string]string{
		"min": "0",
		"max": "2700",
	}

	return []*output.Reading{
		output.Direction.MakeReading("forward"),
		speed,
	}, nil
}

// fanMultiRead is the read handler for the emulated fan-multi device(s).
func fanMultiRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	v := emitter.Next()
	setValue, ok := v.(int)
	if !ok {
		return nil, fmt.Errorf("failed to convert previously set value to integer")
	}

	// Current speed emulates the actual speed of the fan.
	currentSpeed := output.RPM.MakeReading(
		utils.RandIntInRange(setValue-10, setValue+10),
	).WithContext(map[string]string{
		"min": "0",
		"max": "2700",
	})

	// Set speed emulates the value which the fan was set to spin at, which
	// may be different than the current speed as it ramps up/down.
	setSpeed := output.RPM.MakeReading(setValue)

	return []*output.Reading{
		output.Direction.MakeReading("forward"),
		currentSpeed,
		setSpeed,
	}, nil
}

// fanWrite is the write handler for the emulated fan device(s).
func fanWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	// Parse the data []byte into an int
	v, err := strconv.Atoi(string(data.Data))
	if err != nil {
		return err
	}

	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case "speed", "speed_percent":
		emitter.Set(v)
	default:
		return fmt.Errorf("unsupport write action: %v", data.Action)
	}
	return nil
}

// fanMultiWrite is the write handler for the emulated fan-multi device(s).
func fanMultiWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	// Parse the data []byte into an int
	v, err := strconv.Atoi(string(data.Data))
	if err != nil {
		return err
	}

	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case "speed", "speed_percent":
		emitter.Set(v)
	default:
		return fmt.Errorf("unsupport write action: %v", data.Action)
	}
	return nil
}
