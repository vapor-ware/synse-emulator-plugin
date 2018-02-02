package devices

import (
	"fmt"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	stateOn     = "on"
	stateOff    = "off"
	blinkBlink  = "blink"
	blinkSteady = "steady"
)

var (
	state string
	color string
	blink string
)

// EmulatedLED is the handler for the Emulated LED device.
var EmulatedLED = sdk.DeviceHandler{
	Type:  "led",
	Model: "emul8-led",
	Read:  ledRead,
	Write: ledWrite,
}

func ledRead(device *sdk.Device) ([]*sdk.Reading, error) {

	if state == "" {
		state = stateOff
	}
	if color == "" {
		color = "000000"
	}
	if blink == "" {
		blink = blinkSteady
	}

	now := time.Now().String()

	ret := []*sdk.Reading{
		{now, "state", state},
		{now, "color", color},
		{now, "blink", blink},
	}

	return ret, nil
}

func ledWrite(in *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "color" {
		color = string(raw[0])

	} else if action == "blink" {
		cmd := string(raw[0])
		if cmd == blinkSteady {
			blink = blinkSteady
		} else if cmd == blinkBlink {
			blink = blinkBlink
		} else {
			return fmt.Errorf("unsupported command for blink action: %v", cmd)
		}

	} else if action == "state" {
		cmd := string(raw[0])
		if cmd == stateOn {
			state = stateOn
		} else if cmd == stateOff {
			state = stateOff
		} else {
			return fmt.Errorf("unsupported command for state action: %v", cmd)
		}
	}

	return nil
}
