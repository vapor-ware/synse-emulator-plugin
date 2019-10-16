package devices

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	statusStopped  = "stopped"
	statusRotating = "rotating"

	stateReady       = "ok"
	stateUnavailable = "unavailable"
)

// Carousel is the handler for emulated carousel device(s).
var Carousel = sdk.DeviceHandler{
	Name:  "carousel",
	Read:  carouselRead,
	Write: carouselWrite,
}

// carouselRead is the read handler for emulated carousel device(s). It
// returns the state, status, and position values for the device.
func carouselRead(device *sdk.Device) ([]*sdk.Reading, error) {

	// Set default reading values - these will be used if they are not cached.
	state := stateReady
	status := statusStopped
	position := 1

	dState, ok := deviceState[device.GUID()]
	if ok {
		if _, ok := dState["state"]; ok {
			state = dState["state"].(string)
		}

		if _, ok := dState["status"]; ok {
			status = dState["status"].(string)
		}

		if _, ok := dState["position"]; ok {
			position = dState["position"].(int)
		}

	}

	stateReading, err := device.GetOutput("state").MakeReading(state)
	if err != nil {
		return nil, err
	}

	statusReading, err := device.GetOutput("status").MakeReading(status)
	if err != nil {
		return nil, err
	}

	positionReading, err := device.GetOutput("position").MakeReading(position)
	if err != nil {
		return nil, err
	}

	return []*sdk.Reading{
		stateReading,
		statusReading,
		positionReading,
	}, nil
}

// carouselWrite is the write handler for emulated carousel device(s). It
// sets the position value for the device.
func carouselWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Data

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'data', but required")
	}

	if action == "position" {
		position, err := strconv.Atoi(string(raw))
		if err != nil {
			return err
		}

		dState, ok := deviceState[device.GUID()]
		if !ok {
			deviceState[device.GUID()] = map[string]interface{}{"state": stateUnavailable}
			deviceState[device.GUID()] = map[string]interface{}{"status": statusRotating}
		} else {
			dState["state"] = stateUnavailable
			dState["status"] = statusRotating
		}

		go func() {
			// Sleep for a short time before updating the position and changing the
			// state and status back.
			time.Sleep(5 * time.Second)

			dState, ok := deviceState[device.GUID()]
			if !ok {
				deviceState[device.GUID()] = map[string]interface{}{"state": stateReady}
				deviceState[device.GUID()] = map[string]interface{}{"status": statusStopped}
				deviceState[device.GUID()] = map[string]interface{}{"position": position}
			} else {
				dState["state"] = stateReady
				dState["status"] = statusStopped
				dState["position"] = position
			}
		}()
	}
	return nil
}
