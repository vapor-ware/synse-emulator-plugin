package utils

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
)

// deviceEmitters maps the ID of each device to their ValueEmitter.
var deviceEmitters = map[string]*ValueEmitter{}

// GetEmitter gets the ValueEmitter for the specified device.
func GetEmitter(id string) *ValueEmitter {
	if emitter := deviceEmitters[id]; emitter == nil {
		panic(heredoc.Doc(fmt.Sprintf(`
			No value emitter is configured for device %s.

			Ensure that the plugin device configuration is correct and device types
			match those specified in the filters for plugin setup actions.
			(synse-emulator-plugin/pkg/actions.go)`, id,
		)))
	} else {
		return emitter
	}
}

// SetEmitter adds a ValueEmitter for the specified device.
func SetEmitter(id string, emitter *ValueEmitter) error {
	if _, exists := deviceEmitters[id]; exists {
		return fmt.Errorf("emitter already exists for device %s", id)
	}
	deviceEmitters[id] = emitter
	return nil
}
