package outputs

import "github.com/vapor-ware/synse-sdk/sdk/output"

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

	// Position is an output which describes positional state.
	Position = output.Output{
		Name: "position",
		Type: "position",
	}
)
