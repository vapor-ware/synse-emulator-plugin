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

	Counter = output.Output{
		Name:      "openmetrics_counter",
		Type:      "openmetrics_counter",
		Precision: 1,
		Unit:      &output.Unit{},
	}

	Gauge = output.Output{
		Name:      "openmetrics_gauge",
		Type:      "openmetrics_gauge",
		Precision: 1,
		Unit:      &output.Unit{},
	}
)
