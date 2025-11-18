package models

import "time"

type Voucher struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	CrewName     string
	FlightNumber string
	FlightDate   string
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    time.Time
}
