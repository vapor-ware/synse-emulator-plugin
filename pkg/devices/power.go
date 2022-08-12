package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
)

// Power is the handler for the emulated power device(s).
var Power = sdk.DeviceHandler{
	Name:  "power",
	Read:  powerRead,
	Write: minMaxCurrentWrite,
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
