package models

import "time"

type ElectricityReading struct {
	Time    time.Time
	Reading float64
}
