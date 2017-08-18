package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/shogo82148/goa-adder/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}
