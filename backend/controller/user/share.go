package user

import (
	"net/http"
	"pchat/controller/core"
	"pchat/permissions"
)

var Group = core.NewGroup("/users")

func init() {
	Group.Register(core.NewController("/approve", http.MethodPost, approveRegister, core.WithPermission(permissions.REGISTER_APPLICATION_APPROVE)))
	Group.Register(core.NewController("/login", http.MethodPost, login, core.WithNoAuth()))
	Group.Register(core.NewController("/register", http.MethodPost, register, core.WithNoAuth()))
	Group.Register(core.NewController("/:id", http.MethodPut, updateProfile))
	Group.Register(core.NewController("/registerApplications", http.MethodGet, listRegisterApplications, core.WithPermission(permissions.REGISTER_APPLICATION_VIEW)))
	Group.Register(core.NewController("/validOTP", http.MethodPost, validOTP, core.WithAllowUnauthorizedToken()))
	Group.Register(core.NewController("/renewRecoveryCodes", http.MethodPost, renewRecoveryCodes))
	Group.Register(core.NewController("/enable2FA", http.MethodPost, enable2FA))
	Group.Register(core.NewController("/disable2FA", http.MethodPost, disable2FA))
}
