package routes

import (
	"form/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/steps", controllers.GetSteps)
	r.GET("/fields", controllers.GetFields)
	r.GET("/types", controllers.GetTypes)

	r.POST("/kyc/save", controllers.SaveKycData)
	r.POST("/payment/save", controllers.SavePayment)

	r.POST("/identity", controllers.SaveIdentity)
	r.POST("/address", controllers.SaveAddress)
	r.POST("/bank", controllers.SaveBank)
	r.POST("/promoter", controllers.SavePromoters)
	r.POST("/trading", controllers.SaveTrading)
	r.POST("/additional", controllers.SaveAdditional)
}
