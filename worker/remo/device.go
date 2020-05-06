package remo

import "time"

// Device represents a device such as Nature Remo and Nature Remo Mini.
type Device struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	TemperatureOffset int64        `json:"temperature_offset"`
	HumidityOffset    int64        `json:"humidity_offset"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	FirmwareVersion   string       `json:"firmware_version"`
	NewestEvents      NewestEvents `json:"newest_events"`
}

// NewestEvents is the reference key to SensorValue means "te" = temperature, "hu" = humidity, "il" = illumination, "mo" = movement.
type NewestEvents struct {
	Temperature  SensorValue `json:"te" firestore:"te"`
	Humidity     SensorValue `json:"hu" firestore:"hu"`
	Illumination SensorValue `json:"il" firestore:"il"`
	Movement     SensorValue `json:"mo" firestore:"mo"`
}

// SensorValue represents value of sensor.
type SensorValue struct {
	Value     float64   `json:"val" firestore:"val"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}
