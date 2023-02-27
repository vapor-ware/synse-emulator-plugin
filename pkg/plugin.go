package pkg

import (
	"log"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/devices"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
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
		&outputs.JSONOutput,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		&devices.Airflow,
		&devices.Canary,
		&devices.CarouselJSON,
		&devices.CarouselStatus,
		&devices.Current,
		&devices.Energy,
		&devices.Fan,
		&devices.FanMulti,
		&devices.Frequency,
		&devices.Humidity,
		&devices.LED,
		&devices.Lock,
		&devices.UPS,
		&devices.Power,
		&devices.Pressure,
		&devices.Temperature,
		&devices.Voltage,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device setup actions. These will create the emulated
	// devices' value emitters for each device.
	err = plugin.RegisterDeviceSetupActions(
		&ActionAirflowValueEmitterSetup,
		&ActionCanaryValueEmitterSetup,
		&ActionCarouselStatusValueEmitterSetup,
		&ActionCarouselJSONValueEmitterSetup,
		&ActionCurrentValueEmitterSetup,
		&ActionEnergyValueEmitterSetup,
		&ActionFanValueEmitterSetup,
		&ActionFanMultiValueEmitterSetup,
		&ActionFrequencyValueEmitterSetup,
		&ActionHumidityValueEmitterSetup,
		&ActionLEDValueEmitterSetup,
		&ActionLockValueEmitterSetup,
		&ActionUpsDurationValueEmitterSetup,
		&ActionPowerValueEmitterSetup,
		&ActionPressureValueEmitterSetup,
		&ActionTemperatureValueEmitterSetup,
		&ActionVoltageValueEmitterSetup,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
