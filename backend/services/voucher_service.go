package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"user-service/dao"
	"user-service/dto"
	customerror "user-service/errors"
	"user-service/repositories"

	"gorm.io/gorm"
)

type VoucherService interface {
	Generate(req dto.VoucherGenerateReqDTO) (dto.VoucherGenerateRespDTO, error)
	Check(req dto.VoucherCheckReqDTO) (dto.VoucherCheckRespDTO, error)
}

type voucherService struct {
	repo repositories.VoucherRepository
}

// Check implements VoucherService.
func (svc *voucherService) Check(req dto.VoucherCheckReqDTO) (resp dto.VoucherCheckRespDTO, svcErr error) {
	if req.FlightNumber == "" {
		svcErr = customerror.BadRequest("flightNumber is empty")
		return
	}

	if req.Date == "" {
		svcErr = customerror.BadRequest("date is empty")
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)

	if err != nil {
		svcErr = customerror.BadRequest("date format is not valid")
		return
	}

	today := time.Now().Truncate(24 * time.Hour)
	if parsedDate.Before(today) {
		svcErr = customerror.BadRequest("date must be today or later")
		return
	}

	_, err = svc.repo.FindByFlightNumberAndDate(req.FlightNumber, req.Date)

	resp.Exists = true

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Exists = false
		} else {
			svcErr = customerror.Internal(err.Error())
			return
		}
	}

	return
}

// Generate implements VoucherService.
func (svc *voucherService) Generate(req dto.VoucherGenerateReqDTO) (resp dto.VoucherGenerateRespDTO, svcErr error) {

	type AircraftType string

	type SeatRange struct {
		Min    int
		Max    int
		Letter []string
	}

	const (
		ATR          AircraftType = "ATR"
		Airbus320    AircraftType = "Airbus 320"
		Boeing737Max AircraftType = "Boeing 737 Max"
	)

	var AircraftSeatConfig = map[AircraftType]SeatRange{
		ATR: {
			Min:    1,
			Max:    18,
			Letter: []string{"A", "C", "D", "F"},
		},
		Airbus320: {
			Min:    1,
			Max:    32,
			Letter: []string{"A", "B", "C", "D", "E", "F"},
		},
		Boeing737Max: {
			Min:    1,
			Max:    32,
			Letter: []string{"A", "B", "C", "D", "E", "F"},
		},
	}

	if req.ID == "" {
		svcErr = customerror.BadRequest("Crew ID is empty")
		return
	}

	if req.Name == "" {
		svcErr = customerror.BadRequest("Crew Name is empty")
		return
	}

	if req.FlightNumber == "" {
		svcErr = customerror.BadRequest("flightNumber is empty")
		return
	}

	if req.AirCraft == "" {
		svcErr = customerror.BadRequest("aircraft is empty")
		return
	}

	var isValidAirCraft bool

	switch AircraftType(req.AirCraft) {
	case ATR, Airbus320, Boeing737Max:
		isValidAirCraft = true
	default:
		isValidAirCraft = false
	}

	if !isValidAirCraft {
		svcErr = customerror.BadRequest("aircraft type input is not valid")
		return
	}

	if req.Date == "" {
		svcErr = customerror.BadRequest("date is empty")
		return
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)

	if err != nil {
		svcErr = customerror.BadRequest("date format is not valid")
		return
	}

	today := time.Now().Truncate(24 * time.Hour)
	if parsedDate.Before(today) {
		svcErr = customerror.BadRequest("date must be today or later")
		return
	}

	aircraftCfg, ok := AircraftSeatConfig[AircraftType(req.AirCraft)]
	if !ok {
		svcErr = customerror.BadRequest("Aircraft Type input is not valid")
		return
	}

	checkResp, err := svc.Check(dto.VoucherCheckReqDTO{
		FlightNumber: req.FlightNumber,
		Date:         req.Date,
	})

	if err != nil {
		return
	}

	if checkResp.Exists {
		svcErr = customerror.BadRequest("voucher for flightNumber and date inputed has been generated")
		return
	}

	var generatedSeats []string

	for i := 0; i < 3; i++ {
		row := rand.Intn(aircraftCfg.Max-aircraftCfg.Min+1) + aircraftCfg.Min
		col := aircraftCfg.Letter[rand.Intn(len(aircraftCfg.Letter))]
		seat := fmt.Sprintf("%d%s", row, col)
		generatedSeats = append(generatedSeats, seat)
	}

	voucherDao := dao.VoucherDAO{
		CrewName:     req.Name,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.AirCraft,
		Seat1:        generatedSeats[0],
		Seat2:        generatedSeats[1],
		Seat3:        generatedSeats[2],
	}

	err = svc.repo.Create(voucherDao)

	if err != nil {
		svcErr = customerror.Internal(err.Error())
		return
	}

	resp.Success = true
	resp.Seats = generatedSeats

	return
}

func NewVoucherService(r repositories.VoucherRepository) VoucherService {
	return &voucherService{repo: r}
}
