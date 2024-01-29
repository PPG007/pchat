package user

import (
	"net/http"
	"pchat/controller/core"
)

var Group = core.NewGroup("/users")

func init() {
	Group.Register(core.NewController("/approve", http.MethodPost, approveRegister))
	Group.Register(core.NewController("/login", http.MethodPost, login))
	Group.Register(core.NewController("/register", http.MethodPost, register))
	Group.Register(core.NewController("/:id", http.MethodPut, updateProfile))
	Group.Register(core.NewController("/registerApplications", http.MethodGet, listRegisterApplications))
	Group.Register(core.NewController("/validOTP", http.MethodPost, validOTP))
	Group.Register(core.NewController("/renewRecoveryCodes", http.MethodPost, renewRecoveryCodes))
	Group.Register(core.NewController("/enable2FA", http.MethodPost, enable2FA))
	Group.Register(core.NewController("/disable2FA", http.MethodPost, disable2FA))
}
