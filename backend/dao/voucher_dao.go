package dao

import "time"

type VoucherDAO struct {
	ID           uint
	CrewName     string
	FlightNumber string
	FlightDate   string
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    time.Time
}
