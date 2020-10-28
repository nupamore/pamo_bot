package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Controller : controller
type Controller struct{}

// PageMeta : page meta
type PageMeta struct {
	Size int `json:"size"`
	Page int `json:"page"`
	All  int `json:"all"`
}

// Response : response model
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	*PageMeta `json:"pageMeta,omitempty"`
}

// SendError : send error response
func (ctrl *Controller) SendError(c *fiber.Ctx, code int, err error) error {
	res := Response{
		Code:    code,
		Message: err.Error(),
	}
	return c.JSON(res)
}

// error codes
const (
	AuthError         = 9000
	InvalidParamError = 9001
	DBError           = 9002
	EtcError          = 9999
)
