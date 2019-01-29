// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Matsuya-Web-API": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/makotia/Matsuya-Web-API/design
// --out=$(GOPATH)/src/github.com/makotia/Matsuya-Web-API
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
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

// V1Controller is the controller interface for the V1 actions.
type V1Controller interface {
	goa.Muxer
	Random(*RandomV1Context) error
}

// MountV1Controller "mounts" a V1 resource controller on the given service.
func MountV1Controller(service *goa.Service, ctrl V1Controller) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRandomV1Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Random(rctx)
	}
	service.Mux.Handle("GET", "/v1/random", ctrl.MuxHandler("random", h, nil))
	service.LogInfo("mount", "ctrl", "V1", "action", "Random", "route", "GET /v1/random")
}

// V2Controller is the controller interface for the V2 actions.
type V2Controller interface {
	goa.Muxer
	Random(*RandomV2Context) error
}

// MountV2Controller "mounts" a V2 resource controller on the given service.
func MountV2Controller(service *goa.Service, ctrl V2Controller) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRandomV2Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Random(rctx)
	}
	service.Mux.Handle("GET", "/v2/random", ctrl.MuxHandler("random", h, nil))
	service.LogInfo("mount", "ctrl", "V2", "action", "Random", "route", "GET /v2/random")
}

// V3Controller is the controller interface for the V3 actions.
type V3Controller interface {
	goa.Muxer
	Random(*RandomV3Context) error
}

// MountV3Controller "mounts" a V3 resource controller on the given service.
func MountV3Controller(service *goa.Service, ctrl V3Controller) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRandomV3Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Random(rctx)
	}
	service.Mux.Handle("GET", "/v3/random", ctrl.MuxHandler("random", h, nil))
	service.LogInfo("mount", "ctrl", "V3", "action", "Random", "route", "GET /v3/random")
}

// V4Controller is the controller interface for the V4 actions.
type V4Controller interface {
	goa.Muxer
	Random(*RandomV4Context) error
	Search(*SearchV4Context) error
}

// MountV4Controller "mounts" a V4 resource controller on the given service.
func MountV4Controller(service *goa.Service, ctrl V4Controller) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRandomV4Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Random(rctx)
	}
	service.Mux.Handle("GET", "/v4/random", ctrl.MuxHandler("random", h, nil))
	service.LogInfo("mount", "ctrl", "V4", "action", "Random", "route", "GET /v4/random")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSearchV4Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Search(rctx)
	}
	service.Mux.Handle("GET", "/v4/search", ctrl.MuxHandler("search", h, nil))
	service.LogInfo("mount", "ctrl", "V4", "action", "Search", "route", "GET /v4/search")
}
