package outputs

import "github.com/vapor-ware/synse-sdk/v2/sdk/output"

var (
	// Airflow is the output type for airflow readings.
	Airflow = output.Output{
		Name:      "airflow",
		Type:      "speed",
		Precision: 3,
		Unit: &output.Unit{
			Name:   "millimeters per second",
			Symbol: "mm/s",
		},
	}

	// JSONOutput is for arbitrary json.
	JSONOutput = output.Output{
		Name: "json",
		Type: "json",
	}

	// UPS is the output type for duration readings.
	UPS = output.Output{
		Name: "duration",
		Type: "duration",
		Unit: &output.Unit{
			Name:   "seconds",
			Symbol: "s",
		},
	}
)
