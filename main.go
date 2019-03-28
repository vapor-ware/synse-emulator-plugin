package main

import (
	"log"

	"github.com/vapor-ware/synse-emulator-plugin/pkg"
	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	pluginName       = "emulator plugin"
	pluginMaintainer = "vaporio"
	pluginDesc       = "A plugin with emulated devices and data"
	pluginVcs        = "github.com/vapor-ware/synse-emulator-plugin"
)

func main() {
	// Set the plugin metadata
	sdk.SetPluginInfo(
		pluginName,
		pluginMaintainer,
		pluginDesc,
		pluginVcs,
	)

	plugin := pkg.MakePlugin()

	// Run the plugin
	if err := plugin.Run(); err != nil {
		log.Fatal(err)
	}
}
