package outputs

import "github.com/vapor-ware/synse-sdk/sdk"

var (
	// Airflow is the output type for airflow readings.
	Airflow = sdk.OutputType{
		Name:      "airflow",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "millimeters per second",
			Symbol: "mm/s",
		},
	}

	// FanSpeed is the output type for fan speed readings.
	FanSpeed = sdk.OutputType{
		Name:      "fan.speed",
		Precision: 1,
		Unit: sdk.Unit{
			Name:   "revolutions per minute",
			Symbol: "RPM",
		},
	}

	// Humidity is the output type for humidity readings.
	Humidity = sdk.OutputType{
		Name:      "humidity",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "percent humidity",
			Symbol: "%",
		},
	}

	// Temperature is the output type for temperature readings.
	Temperature = sdk.OutputType{
		Name:      "temperature",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "celsius",
			Symbol: "C",
		},
	}

	// LedState is the output type for LED state readings (on/off/blink).
	LedState = sdk.OutputType{
		Name: "led.state",
	}

	// LedColor is the output type for LED color readings.
	LedColor = sdk.OutputType{
		Name: "led.color",
	}

	// Pressure is the output type for differential pressure readings.
	Pressure = sdk.OutputType{
		Name:      "pressure",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "pascals",
			Symbol: "Pa",
		},
	}

	// LockStatus is the output type for Lock status readings (locked/unlocked_electrically).
	LockStatus = sdk.OutputType{
		Name: "lock.status",
	}

	// PowerWatts is the output type for power devices in Watts.
	PowerWatts = sdk.OutputType{
		Name:      "power",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "watt",
			Symbol: "W",
		},
	}
	// Voltage is the output type for voltage devices in voltage.
	Voltage = sdk.OutputType{
		Name:      "voltage",
		Precision: 3,
		Unit: sdk.Unit{
			Name:   "volt",
			Symbol: "V",
		},
	}

	// KWH is the output type for voltage devices in kwh.power.
	KWH = sdk.OutputType{
		Name:      "kwh.power",
		Precision: 5,
		Unit: sdk.Unit{
			Name:   "kilowatt hour",
			Symbol: "kWh",
		},
	}

	// Hertz is the output type for frequency devices in hertz.
	Hertz = sdk.OutputType{
		Name:      "frequency",
		Precision: 4,
		Unit: sdk.Unit{
			Name:   "hertz",
			Symbol: "Hz",
		},
	}

	// Amperage is the output type for amperage devices in ampere.
	Amperage = sdk.OutputType{
		Name:      "current",
		Precision: 2,
		Unit: sdk.Unit{
			Name:   "ampere",
			Symbol: "A",
		},
	}

	// Position is the output type for a positional state.
	Position = sdk.OutputType{
		Name: "position",
	}

	// State is the output type for a device state.
	State = sdk.OutputType{
		Name: "state",
	}

	// Status is the output type for a device status.
	Status = sdk.OutputType{
		Name: "status",
	}
)
