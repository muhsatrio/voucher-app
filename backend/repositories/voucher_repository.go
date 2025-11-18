package repositories

import (
	"user-service/dao"
	"user-service/database"
	"user-service/models"
)

type VoucherRepository interface {
	FindByFlightNumberAndDate(flightNumber string, date string) (dao.VoucherDAO, error)
	Create(data dao.VoucherDAO) error
}

type voucherrepo struct{}

// Create implements VoucherRepository.
func (v *voucherrepo) Create(data dao.VoucherDAO) error {
	newVoucher := models.Voucher{
		CrewName:     data.CrewName,
		FlightNumber: data.FlightNumber,
		FlightDate:   data.FlightDate,
		AircraftType: data.AircraftType,
		Seat1:        data.Seat1,
		Seat2:        data.Seat2,
		Seat3:        data.Seat3,
	}
	err := database.DB.Create(&newVoucher).Error

	return err
}

// FindByFlightNumber implements VoucherRepository.
func (v *voucherrepo) FindByFlightNumberAndDate(flightNumber string, date string) (dao.VoucherDAO, error) {
	var voucher models.Voucher
	err := database.DB.
		Where("flight_number = ? AND flight_date = ?", flightNumber, date).First(&voucher).Error

	if err != nil {
		return dao.VoucherDAO{}, err
	}

	result := dao.VoucherDAO{
		ID:           voucher.ID,
		CrewName:     voucher.CrewName,
		FlightNumber: voucher.FlightNumber,
		FlightDate:   voucher.FlightDate,
		AircraftType: voucher.AircraftType,
		Seat1:        voucher.Seat1,
		Seat2:        voucher.Seat2,
		Seat3:        voucher.Seat3,
		CreatedAt:    voucher.CreatedAt,
	}

	return result, nil
}

func NewVoucherRepository() VoucherRepository {
	return &voucherrepo{}
}
