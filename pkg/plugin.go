package pkg

import (
	"log"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/devices"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the Synse Emulator Plugin.
func MakePlugin() *sdk.Plugin {
	plugin := sdk.NewPlugin()

	// Register the output types
	err := plugin.RegisterOutputTypes(
		&outputs.Airflow,
		&outputs.FanSpeed,
		&outputs.Humidity,
		&outputs.LedColor,
		&outputs.LedState,
		&outputs.Pressure,
		&outputs.Temperature,
		&outputs.LockStatus,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	plugin.RegisterDeviceHandlers(
		&devices.Airflow,
		&devices.Fan,
		&devices.Humidity,
		&devices.LED,
		&devices.Pressure,
		&devices.Temperature,
		&devices.Lock,
	)

	return plugin
}
