package pkg

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/devices"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"log"
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

	handlers := devices.JunosDeviceHandlers()
	handlers = append(handlers,
		&devices.Airflow,
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

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		handlers...,
	)
	if err != nil {
		log.Fatal(err)
	}

	setupActions := OpenMetricsEmitterSetup()
	setupActions = append(setupActions,
		&ActionAirflowValueEmitterSetup,
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

	// Register device setup actions. These will create the emulated
	// devices' value emitters for each device.
	err = plugin.RegisterDeviceSetupActions(
		setupActions...,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
