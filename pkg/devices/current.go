package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

// Current is the handler for the emulated current device(s).
var Current = sdk.DeviceHandler{
	Name:  "current",
	Read:  currentRead,
	Write: minMaxCurrentWrite,
}

// currentRead is the read handler for the emulated current device(s).
func currentRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	ec, err := output.ElectricCurrent.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		ec,
	}, nil
}
