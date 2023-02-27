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
	watt, err := output.Watt.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		watt,
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

	// Parse the data []byte into an int
	v, err := strconv.Atoi(string(data.Data))
	if err != nil {
		return err
	}

	emitter := utils.GetEmitter(device.GetID())
	switch data.Action {
	case MIN:
		emitter.WithLowerBound(v)
	case MAX:
		emitter.WithUpperBound(v)
	case CURRENT:
		emitter.Set(v)
	case STATE:
		emitter.Set(v)
	default:
		return fmt.Errorf("unsupported write action: %v", data.Action)
	}
	return nil
}
