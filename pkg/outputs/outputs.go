package outputs

import "github.com/vapor-ware/synse-sdk/sdk/output"

var (
	// Airflow is the output type for airflow readings.
	Airflow = output.Output{
		Name:      "airflow",
		Precision: 3,
		Unit: &output.Unit{
			Name:   "millimeters per second",
			Symbol: "mm/s",
		},
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
)
