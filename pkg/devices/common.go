package devices

import (
	"fmt"
	"strconv"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
)

// Constants for common device write actions.
const (
	MIN     = "min"
	MAX     = "max"
	CURRENT = "current"
)

// minMaxCurrentWrite is a generic write handler which can be used for devices
// which can not typically be written to. In the emulator, writing is enabled on
// devices where it would not otherwise be possible in order to adjust the reading
// values and ranges at runtime without having to re-deploy.
func minMaxCurrentWrite(device *sdk.Device, data *sdk.WriteData) error {
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
	default:
		return fmt.Errorf("unsupported write action: %v", data.Action)
	}
	return nil
}
