package controllers

import (
	"github.com/goadesign/goa"
	"github.com/waymobetta/valhalla-api/app"
	"github.com/waymobetta/valhalla-api/db"
)

// GodkjentController implements the godkjent resource.
type GodkjentController struct {
	*goa.Controller
}

// NewGodkjentController creates a godkjent controller.
func NewGodkjentController(service *goa.Service) *GodkjentController {
	return &GodkjentController{
		Controller: service.NewController("GodkjentController"),
	}
}

// LeggeTil runs the leggeTil action.
func (c *GodkjentController) LeggeTil(ctx *app.LeggeTilGodkjentContext) error {
	// GodkjentController_LeggeTil: start_implement

	// Put your logic here

	secret := ctx.Payload.Secret
	if secret != "raido" {
		err := &app.StandardError{
			Code:    500,
			Message: "could not authenticate",
		}
		return ctx.InternalServerError(err)
	}

	err := db.LeggeTil(
		ctx.Payload.Navn,
		ctx.Payload.Adresse,
	)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res := &app.Godkjent{
		Navn:    ctx.Payload.Navn,
		Adresse: ctx.Payload.Adresse,
	}

	return ctx.OK(res)
	// GodkjentController_LeggeTil: end_implement
}

// Vis runs the vis action.
func (c *GodkjentController) Vis(ctx *app.VisGodkjentContext) error {
	// GodkjentController_Vis: start_implement

	// Put your logic here

	godkjentListe, err := db.Vis()
	if err != nil {
		return ctx.InternalServerError(err)
	}

	res := &app.Godkjentliste{
		GodkjentListe: godkjentListe,
	}
	return ctx.OK(res)
	// GodkjentController_Vis: end_implement
}
