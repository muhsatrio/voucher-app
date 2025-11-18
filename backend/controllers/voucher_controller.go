package controllers

import (
	"net/http"
	"user-service/dto"
	customerror "user-service/errors"
	"user-service/services"

	"github.com/gin-gonic/gin"
)

type VoucherController struct {
	service services.VoucherService
}

func NewVoucherController(s services.VoucherService) *VoucherController {
	return &VoucherController{service: s}
}

func (ctrl *VoucherController) VoucherCheck(c *gin.Context) {
	type voucherCheckReq struct {
		FlightNumber string `json:"flightNumber"`
		Date         string `json:"date"`
	}

	type voucherCheckResp struct {
		Exists bool `json:"exists"`
	}

	var req voucherCheckReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"error":  err.Error(),
		})
		return
	}

	resp, err := ctrl.service.Check(dto.VoucherCheckReqDTO(req))

	if err != nil {
		if svcErr, ok := err.(*customerror.ServiceError); ok {
			c.JSON(svcErr.Code, gin.H{
				"error": svcErr.Message,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	respController := voucherCheckResp{
		Exists: resp.Exists,
	}

	c.JSON(http.StatusOK, respController)
}

func (ctrl *VoucherController) VoucherGenerate(c *gin.Context) {
	type voucherGenerateReq struct {
		Name         string `json:"name"`
		ID           string `json:"id"`
		FlightNumber string `json:"flightNumber"`
		Date         string `json:"date"`
		AirCraft     string `json:"aircraft"`
	}

	type voucherGenerateResp struct {
		Success bool     `json:"success"`
		Seats   []string `json:"seats"`
	}

	var req voucherGenerateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": false,
			"error":  err.Error(),
		})
		return
	}

	resp, err := ctrl.service.Generate(dto.VoucherGenerateReqDTO(req))

	if err != nil {
		if svcErr, ok := err.(*customerror.ServiceError); ok {
			c.JSON(svcErr.Code, gin.H{
				"error": svcErr.Message,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	respController := voucherGenerateResp{
		Success: resp.Success,
		Seats:   resp.Seats,
	}

	c.JSON(http.StatusOK, respController)
}
