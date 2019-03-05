package devices

// deviceState tracks state for devices. The map key is the device id. The value
// map maps arbitrary keys to arbitrary values.
var deviceState = map[string]map[string]interface{}{}

// Constants for common device state fields.
const (
	MIN = "min"
	MAX = "max"
)
