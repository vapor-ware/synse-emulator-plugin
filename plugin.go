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
// through its device configuration.
func ProtocolIdentifier(data map[string]string) string {
	return data["id"]
}

func main() {
	// Create a new Plugin and configure it.
	plugin := sdk.NewPlugin()
	err := plugin.Configure()
	if err != nil {
		log.Fatal(err)
	}

	plugin.RegisterDeviceIdentifier(ProtocolIdentifier)
	plugin.RegisterDeviceHandlers(
		&devices.EmulatedFan,
		&devices.EmulatedLED,
		&devices.EmulatedTemp,
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
