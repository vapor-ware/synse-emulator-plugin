package devices

import (
	"fmt"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
	"strconv"
)

// Battery is the handler for emulated battery devices.
var Battery = sdk.DeviceHandler{
	Name: "battery",
	Read: batteryRead,
	Write: batteryWrite,
}

func batteryRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		output.Status.MakeReading(emitter.Next()).WithContext(map[string]string{
			"metric_name": device.Info,
		}),
	}, nil
}


func batteryWrite(device *sdk.Device, data *sdk.WriteData) error {
	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case "set":
		var val interface{}
		var err error

		t := device.Data["type"]
		switch t {
		case "int":
			val, err = strconv.Atoi(string(data.Data))
			if err != nil {
				return err
			}
		case nil:
			val = string(data.Data)
		default:
			return fmt.Errorf("unsupported 'type' config set for status device: %v", t)
		}
		emitter.Set(val)

	default:
		return fmt.Errorf("unsupported action for status write: %v", data.Action)
	}
	return nil
}