package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// Controller : controller
type Controller struct{}

// Response : response model
type Response struct {
	Code    int         `json:"code"`
	Message error       `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SendError : send error response
func (ctrl *Controller) SendError(c *fiber.Ctx, code int, err error) error {
	res := Response{
		Code:    code,
		Message: err,
	}
	r, _ := json.Marshal(res)
	return c.Send(r)
}

// error codes
const (
	InvalidParamError = 9001
	DBError           = 9002
	EtcError          = 9999
)
