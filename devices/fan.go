package devices

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vapor-ware/synse-sdk/sdk"
)

var speed int

// EmulatedFan is the handler for the Emulated fan device.
var EmulatedFan = sdk.DeviceHandler{
	Type:  "fan",
	Model: "emul8-fan",
	Read:  fanRead,
	Write: fanWrite,
}

func fanRead(device *sdk.Device) ([]*sdk.Reading, error) {
	now := time.Now().String()
	ret := []*sdk.Reading{
		{now, "fan_speed", strconv.Itoa(speed)},
	}
	return ret, nil
}

func fanWrite(device *sdk.Device, data *sdk.WriteData) error {
	action := data.Action
	raw := data.Raw

	// We always expect the action to come with raw data, so if it
	// doesn't exist, error.
	if len(raw) == 0 {
		return fmt.Errorf("no values specified for 'raw', but required")
	}

	if action == "speed" {
		s, err := strconv.Atoi(string(raw[0]))
		if err != nil {
			return err
		}
		speed = s
	}

	return nil
}
