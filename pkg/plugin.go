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
		&outputs.PowerWatts,
		&outputs.KWH,
		&outputs.Voltage,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		&devices.Airflow,
		&devices.Fan,
		&devices.Humidity,
		&devices.LED,
		&devices.Pressure,
		&devices.Temperature,
		&devices.Lock,
		&devices.Power,
		&devices.Voltage,
		&devices.KillaWattHours,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
