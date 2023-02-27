package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

// Power is the handler for the emulated power device(s).
var Power = sdk.DeviceHandler{
	Name:  "power",
	Read:  powerRead,
	Write: powerWrite,
}

// powerRead is the read handler for the emulated power device(s).
func powerRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	emitterState := utils.GetEmitter(fmt.Sprintf("%s-state", device.GetID()))
	val := emitterState.Next().(map[string]string)

	watt, err := output.Watt.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}

	state, err := output.State.MakeReading(val["state"])
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		watt,
		state,
	}, nil
}

// Constants for common device write actions.
const (
	STATE = "state"
)

func powerWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	emitter := utils.GetEmitter(device.GetID())
	emitterState := utils.GetEmitter(fmt.Sprintf("%s-state", device.GetID()))
	current := emitterState.Next().(map[string]string)

	switch data.Action {
	case MIN:
		v, err := strconv.Atoi(string(data.Data))
		if err != nil {
			return err
		}
		emitter.WithLowerBound(v)
	case MAX:
		v, err := strconv.Atoi(string(data.Data))
		if err != nil {
			return err
		}
		emitter.WithUpperBound(v)
	case CURRENT:
		v, err := strconv.Atoi(string(data.Data))
		if err != nil {
			return err
		}
		emitter.Set(v)
	case STATE:
		current["state"] = string(data.Data)
		emitterState.Set(current)
	default:
		return fmt.Errorf("unsupported write action: %v", data.Action)
	}
	return nil
}
