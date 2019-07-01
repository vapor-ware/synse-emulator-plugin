package devices

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

// Airflow is the handler for the emulated airflow device(s).
var Airflow = sdk.DeviceHandler{
	Name:  "airflow",
	Read:  airflowRead,
	Write: minMaxCurrentWrite,
}

// airflowRead is the read handler for the emulated airflow device(s).
func airflowRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	return []*output.Reading{
		outputs.Airflow.MakeReading(emitter.Next()),
	}, nil
}
