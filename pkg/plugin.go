package pkg

import (
	"log"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/devices"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the Synse Emulator Plugin.
func MakePlugin() *sdk.Plugin {
	plugin, err := sdk.NewPlugin()
	if err != nil {
		log.Fatal(err)
	}

	// Register custom output types.
	err = plugin.RegisterOutputs(
		&outputs.Airflow,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		&devices.Airflow,
		&devices.Energy,
		&devices.Fan,
		&devices.Humidity,
		&devices.LED,
		&devices.Lock,
		&devices.Power,
		&devices.Pressure,
		&devices.Temperature,
		&devices.Voltage,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
