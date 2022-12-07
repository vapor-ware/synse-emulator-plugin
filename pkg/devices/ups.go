package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

// UPS is the handler for the emulated current device(s).
var UPS = sdk.DeviceHandler{
	Name:  "seconds",
	Read:  secondsRead,
	Write: minMaxCurrentWrite,
}

// secondsRead is the read handler for the emulated current device(s).
func secondsRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	ec, err := output.Seconds.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}

	return []*output.Reading{
		ec,
	}, nil
}
