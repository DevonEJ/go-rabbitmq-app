package dto

import (
	"time"
	"encoding/gob"
)

type SensorMessage struct {
	Name string
	Value float64
	Timestamp time.Time
}

// Make the sensor message object ready to send via gob
func init() {
	gob.Register(SensorMessage{})
}
