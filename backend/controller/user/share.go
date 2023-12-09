package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(e *gin.Engine) {
	group := e.Group("/users")
	group.POST("/login", LoginController)
	group.POST("/register", RegisterController)
	group.POST("/approve", ApproveRegisterController)
	group.PUT("/:id", UpdateProfileController)
	group.GET("/registerApplications", ListRegisterApplicationController)
	group.POST("/validOTP", ValidOTP)
	group.POST("/renewRecoveryCodes", RenewRecoveryCodes)
}
