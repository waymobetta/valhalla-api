// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "valhalla": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/waymobetta/valhalla-api/design
// --out=$(GOPATH)/src/github.com/waymobetta/valhalla-api
// --version=v1.4.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// GodkjentController is the controller interface for the Godkjent actions.
type GodkjentController interface {
	goa.Muxer
	LeggeTil(*LeggeTilGodkjentContext) error
	Vis(*VisGodkjentContext) error
}

// MountGodkjentController "mounts" a Godkjent resource controller on the given service.
func MountGodkjentController(service *goa.Service, ctrl GodkjentController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/godkjent", ctrl.MuxHandler("preflight", handleGodkjentOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLeggeTilGodkjentContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LeggeTilGodkjentNyttelast)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.LeggeTil(rctx)
	}
	h = handleGodkjentOrigin(h)
	service.Mux.Handle("POST", "/v1/godkjent", ctrl.MuxHandler("leggeTil", h, unmarshalLeggeTilGodkjentPayload))
	service.LogInfo("mount", "ctrl", "Godkjent", "action", "LeggeTil", "route", "POST /v1/godkjent")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVisGodkjentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Vis(rctx)
	}
	h = handleGodkjentOrigin(h)
	service.Mux.Handle("GET", "/v1/godkjent", ctrl.MuxHandler("vis", h, nil))
	service.LogInfo("mount", "ctrl", "Godkjent", "action", "Vis", "route", "GET /v1/godkjent")
}

// handleGodkjentOrigin applies the CORS response headers corresponding to the origin.
func handleGodkjentOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLeggeTilGodkjentPayload unmarshals the request body into the context request data Payload field.
func unmarshalLeggeTilGodkjentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &leggeTilGodkjentNyttelast{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HelseController is the controller interface for the Helse actions.
type HelseController interface {
	goa.Muxer
	Vis(*VisHelseContext) error
}

// MountHelseController "mounts" a Helse resource controller on the given service.
func MountHelseController(service *goa.Service, ctrl HelseController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/helse", ctrl.MuxHandler("preflight", handleHelseOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVisHelseContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Vis(rctx)
	}
	h = handleHelseOrigin(h)
	service.Mux.Handle("GET", "/v1/helse", ctrl.MuxHandler("vis", h, nil))
	service.LogInfo("mount", "ctrl", "Helse", "action", "Vis", "route", "GET /v1/helse")
}

// handleHelseOrigin applies the CORS response headers corresponding to the origin.
func handleHelseOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, POST, GET, UPDATE, DELETE, PATCH")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
