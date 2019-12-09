package devices

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

const (
	statusStopped  = "stopped"
	statusRotating = "rotating"

	stateReady       = "ready"
	stateUnavailable = "unavailable"
)

// Carousel is the handler for emulator carousel device(s).
var Carousel = sdk.DeviceHandler{
	Name:  "carousel",
	Read:  carouselRead,
	Write: carouselWrite,
}

// carouselRead is the read handler for emulated carousel device(s). It
// returns the state, status, and position values for the device.
func carouselRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	val := emitter.Next().(map[string]interface{})
	return []*output.Reading{
		output.State.MakeReading(val["state"]),
		output.Status.MakeReading(val["status"]),
		outputs.Position.MakeReading(val["position"]),
	}, nil
}

// carouselWrite is the write handler for emulated carousel device(s). It
// sets the position value for the device.
func carouselWrite(device *sdk.Device, data *sdk.WriteData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	emitter := utils.GetEmitter(device.GetID())
	current := emitter.Next().(map[string]interface{})

	switch data.Action {
	case "position":
		pos, err := strconv.Atoi(string(data.Data))
		if err != nil {
			return err
		}

		// Set the state to designate that the carousel is rotating.
		current["state"] = stateUnavailable
		current["status"] = statusRotating

		// After a short while, update the state so the carousel is done
		// rotating and in the new position.
		go func() {
			time.Sleep(5 * time.Second)

			current["state"] = stateReady
			current["status"] = statusStopped
			current["position"] = pos
			emitter.Set(current)
		}()

	default:
		return fmt.Errorf("unsupported write action: %v", data.Action)
	}
	return nil
}
