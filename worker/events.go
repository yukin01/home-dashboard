package main

import (
	"time"

	"github.com/yukin01/home-dashboard/worker/remo"
)

type SensorValue struct {
	Value     float64   `firestore:"val"`
	CreatedAt time.Time `firestore:"created_at"`
}

type Events struct {
	Temperature  SensorValue `firestore:"te"`
	Humidity     SensorValue `firestore:"hu"`
	Illumination SensorValue `firestore:"il"`
	Movement     SensorValue `firestore:"mo"`
	CreatedAt    time.Time   `firestore:"created_at"`
}

func NewSensorValue(s remo.SensorValue) SensorValue {
	return SensorValue{Value: s.Value, CreatedAt: s.CreatedAt}
}

func NewEvents(e remo.NewestEvents) Events {
	return Events{
		Temperature:  NewSensorValue(e.Temperature),
		Humidity:     NewSensorValue(e.Humidity),
		Illumination: NewSensorValue(e.Illumination),
		Movement:     NewSensorValue(e.Movement),
		CreatedAt:    time.Now(),
	}
}
