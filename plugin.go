package main

import (
	"log"

	"github.com/vapor-ware/synse-emulator-plugin/devices"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// Build time variables for setting the version info of a Plugin.
var (
	BuildDate     string
	GitCommit     string
	GitTag        string
	GoVersion     string
	VersionString string
)

// ProtocolIdentifier defines the emulator-specific way of uniquely identifying a device
// through its device configuration. For emulator devices, it expects to find an "id"
// field in the instance configuration and will use that to help construct the unique
// device ID.
func ProtocolIdentifier(data map[string]string) string {
	return data["id"]
}

func main() {
	// Create the protocol identifier for the emulator.
	handlers, err := sdk.NewHandlers(ProtocolIdentifier, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Plugin with the configuration from the default paths.
	plugin, err := sdk.NewPlugin(handlers, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Register the devices with the plugin
	plugin.RegisterDeviceHandlers(
		&devices.EmulatedFan,
		&devices.EmulatedLED,
		&devices.EmulatedTemp,
		&devices.EmulatedAirflow,
		&devices.EmulatedHumidity,
		&devices.EmulatedPressure,
	)

	// Set build-time version info
	plugin.SetVersion(sdk.VersionInfo{
		BuildDate:     BuildDate,
		GitCommit:     GitCommit,
		GitTag:        GitTag,
		GoVersion:     GoVersion,
		VersionString: VersionString,
	})

	// Run the plugin.
	err = plugin.Run()
	if err != nil {
		log.Fatal(err)
	}
}
