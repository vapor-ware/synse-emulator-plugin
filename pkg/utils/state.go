package utils

import "fmt"

// deviceEmitters maps the ID of each device to their ValueEmitter.
var deviceEmitters = map[string]*ValueEmitter{}

// GetEmitter gets the ValueEmitter for the specified device.
func GetEmitter(id string) *ValueEmitter {
	return deviceEmitters[id]
}

// SetEmitter adds a ValueEmitter for the specified device.
func SetEmitter(id string, emitter *ValueEmitter) error {
	if _, exists := deviceEmitters[id]; exists {
		return fmt.Errorf("emitter already exists for device %s", id)
	}
	deviceEmitters[id] = emitter
	return nil
}
