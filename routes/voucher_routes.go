package routes

import (
	"user-service/controllers"
	"user-service/repositories"
	"user-service/services"

	"github.com/gin-gonic/gin"
)

func SetupVoucherRoutes(r *gin.Engine) {
	voucherRepo := repositories.NewVoucherRepository()
	voucherService := services.NewVoucherService(voucherRepo)
	voucherController := controllers.NewVoucherController(voucherService)

	voucherGroup := r.Group("/api")
	{
		voucherGroup.POST("/generate", voucherController.VoucherGenerate)
		voucherGroup.POST("/check", voucherController.VoucherCheck)
	}
}
