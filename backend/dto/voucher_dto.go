package dto

type VoucherCheckReqDTO struct {
	FlightNumber string
	Date         string
}

type VoucherCheckRespDTO struct {
	Exists bool
}

type VoucherGenerateReqDTO struct {
	Name         string
	ID           string
	FlightNumber string
	Date         string
	AirCraft     string
}

type VoucherGenerateRespDTO struct {
	Success bool
	Seats   []string
}
