package devices

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/sdk"
	"github.com/vapor-ware/synse-sdk/sdk/output"
)

const (
	// This is not the full set of possible statuses, but it suffices for a simple emulation.
	statusStopped  = 0
	statusRotating = 1
)

// CarouselStatus is the handler for emulator carousel status device(s).
var CarouselStatus = sdk.DeviceHandler{
	Name: "status",
	Read: carouselRead,
}

// CarouselJSON is the handler for emulator carousel json device(s), which let you "write"
// to the carousel controller.
var CarouselJSON = sdk.DeviceHandler{
	Name:  "json",
	Read:  carouselJSONRead,
	Write: carouselJSONWrite,
}

// carouselRead is the read handler for emulated carousel status device(s). It
// returns the status reading for the device. The default statuses for devices
// are determined based on the device Info field. These defaults are set in the
// plugin's actions.go file, in the startup action which initializes and registers
// the emitter for the device type.
func carouselRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	return []*output.Reading{
		output.Status.MakeReading(emitter.Next()),
	}, nil
}

// carouselJSONRead is the read handler for emulated carousel json device(s). It
// returns whether or not the carousel can be rotated.
func carouselJSONRead(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())

	// The map contains a success value under the "ok" key and an error value
	// under the "error" key. We can check whether we are in error mode or success
	// mode by checking the "mode" key.
	val := emitter.Next().(map[string]string)

	return []*output.Reading{
		outputs.JSONOutput.MakeReading(val[val["mode"]]),
	}, nil
}

// carouselWriteAction models the expected Action that is received by the emulator
// in order to write to the carousel controller.
type carouselWriteAction struct {
	Rack int `json:"rack,omitempty"`
}

// carouseJSONWrite is the write handler for emulated carousel device(s). It
// sets the position which the carousel should be rotated to.
func carouselJSONWrite(device *sdk.Device, data *sdk.WriteData) error {

	emitter := utils.GetEmitter(device.GetID())
	deviceState := emitter.Next().(map[string]string)
	mode := deviceState["mode"]

	setRackEmitter := utils.GetEmitter(utils.CarouselSetRackPositionDevice.GetID())
	getRackEmitter := utils.GetEmitter(utils.CarouselGetRackPositionDevice.GetID())
	stateMachineEmitter := utils.GetEmitter(utils.CarouselStateMachine.GetID())
	statusEmitter := utils.GetEmitter(utils.CarouselStatusDevice.GetID())

	setRackValue := setRackEmitter.Next().(int)

	if data.Action == "mode" {
		switch string(data.Data) {
		case "ok":
			if deviceState["mode"] != "ok" {
				log.Info("setting carousel into success mode")
				deviceState["mode"] = "ok"

				// Put the read-only devices into a success state.
				utils.CarouselMutex.Lock()

				stateMachineEmitter.Set(1)
				statusEmitter.Set(0)

				utils.CarouselMutex.Unlock()
			}

		case "error":
			if deviceState["mode"] != "error" {
				log.Info("setting carousel into error mode")
				deviceState["mode"] = "error"

				// Put the read-only devices into an error state.
				utils.CarouselMutex.Lock()

				stateMachineEmitter.Set(6)
				statusEmitter.Set(1 << 8)

				utils.CarouselMutex.Unlock()
			}

		}
		emitter.Set(deviceState)
		return nil
	}

	// If we are in error mode, return an error.
	if mode == "error" {
		return fmt.Errorf("failed to rotate carousel")
	}

	// Load the write action. This will tell us which "rack" to "rotate to".
	var action carouselWriteAction
	if err := json.Unmarshal([]byte(data.Action), &action); err != nil {
		return err
	}

	if setRackValue == action.Rack {
		// If we told it to rotate to the position it is already at, do nothing.
		return nil
	}

	// Otherwise, we need to set the read-only values accordingly. The SetRack value
	// gets the end state value. The Status will change to the state for "rotating".
	// The GetRack value gets the current rack, which we step through in order to get
	// to the target rack, with a bit of timed interval in between.
	utils.CarouselMutex.Lock()
	setRackEmitter.Set(action.Rack)
	statusEmitter.Set(statusRotating)
	stateMachineEmitter.Set(1)
	go func() {
		var currentPos = setRackValue
		for {
			// Check that the current position is equal to the target position.
			if currentPos == action.Rack {
				break
			}
			time.Sleep(5 * time.Second)

			// Modulo 6 since we have 6 total racks on the carousel.
			currentPos = (currentPos + 1) % 6
			getRackEmitter.Set(currentPos)
		}

		// Once we are done "rotating", set the status back to stopped.
		statusEmitter.Set(statusStopped)
		utils.CarouselMutex.Unlock()
	}()

	return nil
}
