package devices

import (
	"github.com/ghodss/yaml"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/outputs"
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
	"github.com/vapor-ware/synse-sdk/v2/sdk/config"
	"github.com/vapor-ware/synse-sdk/v2/sdk/output"
	"os"
)

const (
	// Config file locations
	localDeviceConfig   = "./config/device/junos.yaml"
	defaultDeviceConfig = "/etc/synse/plugin/config/device/junos.yaml"
)

// JunosDeviceHandlers creates device handlers from a config
func JunosDeviceHandlers() []*sdk.DeviceHandler {
	handlers := []*sdk.DeviceHandler{}
	devices, err := DevicesFromConfig()
	if err != nil {
		return handlers
	}

	for _, d := range devices {
		h := sdk.DeviceHandler{
			Name: d.Type,
		}

		if len(d.Instances) == 0 {
			continue
		}

		if val, ok := d.Instances[0].Data["type"]; ok {
			if val == "Gauge" {
				h.Read = gaugeFunc
			} else if val == "Counter" {
				h.Read = counterFunc
			}
		} else {
			return handlers
		}
		handlers = append(handlers, &h)
	}
	return handlers
}

// DevicesFromConfig parses for junos device config and returns device prototypes
func DevicesFromConfig() ([]*config.DeviceProto, error) {
	path := defaultDeviceConfig
	if _, err := os.Stat(defaultDeviceConfig); err != nil {
		if os.IsNotExist(err) {
			path = localDeviceConfig
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	devices := &config.Devices{}
	if err = yaml.Unmarshal(data, devices); err != nil {
		return nil, err
	}

	if devices.Devices == nil {
		return nil, err
	}
	return devices.Devices, nil
}

func gaugeFunc(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	data, err := outputs.Gauge.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		data,
	}, nil
}

func counterFunc(device *sdk.Device) ([]*output.Reading, error) {
	emitter := utils.GetEmitter(device.GetID())
	data, err := outputs.Counter.MakeReading(emitter.Next())
	if err != nil {
		return nil, err
	}
	return []*output.Reading{
		data,
	}, nil
}
