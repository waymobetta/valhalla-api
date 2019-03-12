package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/valhalla-api/app"
)

// HelseController implements the helse resource.
type HelseController struct {
	*goa.Controller
}

// NewHelseController creates a helse controller.
func NewHelseController(service *goa.Service) *HelseController {
	return &HelseController{
		Controller: service.NewController("HelseController"),
	}
}

// Vis runs the vis action.
func (c *HelseController) Vis(ctx *app.VisHelseContext) error {
	// HelseController_Vis: start_implement

	// Put your logic here

	res := &app.Helse{
		Status: "OK",
	}
	return ctx.OK(res)
	// HelseController_Vis: end_implement
}
