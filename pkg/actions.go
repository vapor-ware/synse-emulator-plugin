package pkg

import (
	"github.com/vapor-ware/synse-emulator-plugin/pkg/utils"
	"github.com/vapor-ware/synse-sdk/v2/sdk"
)

// ActionAirflowValueEmitterSetup initializes a ValueEmitter for each "airflow" type device.
var ActionAirflowValueEmitterSetup = sdk.DeviceAction{
	Name: "airflow value emitter setup",
	Filter: map[string][]string{
		"type": {"airflow"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = -100
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 100
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionCarouselStatusValueEmitterSetup initializes a ValueEmitter for each "status" type device (for carousels).
var ActionCarouselStatusValueEmitterSetup = sdk.DeviceAction{
	Name: "Carousel status value emitter setup",
	Filter: map[string][]string{
		"type": {"status"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		// Define the default values for different carousel status devices
		var seed interface{}

		switch d.Info {
		case "Carousel Status Register":
			seed = 0
			utils.CarouselStatusDevice = d
		case "Carousel VFD Error Code":
			seed = 0
		case "Carousel Get Rack Position":
			seed = 1
			utils.CarouselGetRackPositionDevice = d
		case "Carousel State Machine Code":
			seed = 1
			utils.CarouselStateMachine = d
		case "Carousel Set Rack Position":
			seed = 1
			utils.CarouselSetRackPositionDevice = d
		}

		// Set the default seed to the value determined above, based on the device's
		// Info field.
		emitter := utils.NewValueEmitter(utils.Store).WithSeed(seed)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionCarouselJSONValueEmitterSetup initializes a ValueEmitter for each "json" type device (for carousel).
var ActionCarouselJSONValueEmitterSetup = sdk.DeviceAction{
	Name: "Carousel json value emitter setup",
	Filter: map[string][]string{
		"type": {"json"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed(map[string]string{
			"mode":  "ok",
			"ok":    `{"errors": {}, "status": "ok"}`,
			"error": `{"status":"fail","errors":{"chamber_locks":{"configuration":{"additional":[],"missing":[{"rack_id":"r4","device_info":"L1-36B Lock 1"},{"rack_id":"r4","device_info":"L1-36A Lock 5"},{"rack_id":"r4","device_info":"L1-36F Lock 9"}]}}}}`,
		})
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionCurrentValueEmitterSetup initializes a ValueEmitter for each "current" type device.
var ActionCurrentValueEmitterSetup = sdk.DeviceAction{
	Name: "current value emitter setup",
	Filter: map[string][]string{
		"type": {"current"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 30
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionEnergyValueEmitterSetup initializes a ValueEmitter for each "energy" type device.
var ActionEnergyValueEmitterSetup = sdk.DeviceAction{
	Name: "energy value emitter setup",
	Filter: map[string][]string{
		"type": {"energy"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 100000
		}

		emitter := utils.NewValueEmitter(utils.Accumulate).WithLowerBound(lowerBound).WithUpperBound(upperBound)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionFanValueEmitterSetup initializes a ValueEmitter for each "fan" type device.
var ActionFanValueEmitterSetup = sdk.DeviceAction{
	Name: "fan value emitter setup",
	Filter: map[string][]string{
		"type": {"*fan"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		if d.GetHandler().Name == "fan" {
			seed, ok := d.Data["seed"].(int)
			if !ok {
				seed = 0
			}

			emitter := utils.NewValueEmitter(utils.Store).WithSeed(seed)
			return utils.SetEmitter(d.GetID(), emitter)
		}
		return nil
	},
}

// ActionFanMultiValueEmitterSetup initializes a ValueEmitter for each "fan-multi" type device.
var ActionFanMultiValueEmitterSetup = sdk.DeviceAction{
	Name: "fan-multi value emitter setup",
	Filter: map[string][]string{
		"type": {"*fan"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		if d.GetHandler().Name == "fan-multi" {
			seed, ok := d.Data["seed"].(int)
			if !ok {
				seed = 0
			}

			emitter := utils.NewValueEmitter(utils.Store).WithSeed(seed)
			return utils.SetEmitter(d.GetID(), emitter)
		}
		return nil
	},
}

// ActionFrequencyValueEmitterSetup initializes a ValueEmitter for each "frequency" type device.
var ActionFrequencyValueEmitterSetup = sdk.DeviceAction{
	Name: "frequency value emitter setup",
	Filter: map[string][]string{
		"type": {"frequency"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 60
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionHumidityValueEmitterSetup initializes a ValueEmitter for each "humidity" type device.
var ActionHumidityValueEmitterSetup = sdk.DeviceAction{
	Name: "humidity value emitter setup",
	Filter: map[string][]string{
		"type": {"humidity"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 100
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionLEDValueEmitterSetup initializes a ValueEmitter for each "led" type device.
var ActionLEDValueEmitterSetup = sdk.DeviceAction{
	Name: "LED value emitter setup",
	Filter: map[string][]string{
		"type": {"led"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed(map[string]string{
			"state": "off",
			"color": "000000",
		})
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionLockValueEmitterSetup initializes a ValueEmitter for each "lock" type device.
var ActionLockValueEmitterSetup = sdk.DeviceAction{
	Name: "lock value emitter setup",
	Filter: map[string][]string{
		"type": {"lock"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		emitter := utils.NewValueEmitter(utils.Store).WithSeed("locked")
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionUpsDurationValueEmitterSetup initializes a ValueEmitter for each "seconds" type device
var ActionUpsDurationValueEmitterSetup = sdk.DeviceAction{
	Name: "UPS Duration emitter setup",
	Filter: map[string][]string{
		"type": {"seconds"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 0
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.Accumulate).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionPowerValueEmitterSetup initializes a ValueEmitter for each "power" type device.
var ActionPowerValueEmitterSetup = sdk.DeviceAction{
	Name: "power value emitter setup",
	Filter: map[string][]string{
		"type": {"power"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 1000
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 3000
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionPressureValueEmitterSetup initializes a ValueEmitter for each "pressure" type device.
var ActionPressureValueEmitterSetup = sdk.DeviceAction{
	Name: "pressure value emitter setup",
	Filter: map[string][]string{
		"type": {"pressure"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = -5
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 5
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionTemperatureValueEmitterSetup initializes a ValueEmitter for each "temperature" type device.
var ActionTemperatureValueEmitterSetup = sdk.DeviceAction{
	Name: "temperature value emitter setup",
	Filter: map[string][]string{
		"type": {"temperature"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 0
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 100
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}

// ActionVoltageValueEmitterSetup initializes a ValueEmitter for each "voltage" type device.
var ActionVoltageValueEmitterSetup = sdk.DeviceAction{
	Name: "voltage value emitter setup",
	Filter: map[string][]string{
		"type": {"voltage"},
	},
	Action: func(_ *sdk.Plugin, d *sdk.Device) error {
		lowerBound, ok := d.Data["min"].(int)
		if !ok {
			lowerBound = 100
		}

		upperBound, ok := d.Data["max"].(int)
		if !ok {
			upperBound = 500
		}

		step, ok := d.Data["step"].(int)
		if !ok {
			step = 0
		}

		emitter := utils.NewValueEmitter(utils.RandomWalk).WithLowerBound(lowerBound).WithUpperBound(upperBound).WithStep(step)
		return utils.SetEmitter(d.GetID(), emitter)
	},
}
