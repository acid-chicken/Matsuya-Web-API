package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
)

// V1Controller implements the v1 resource.
type V1Controller struct {
	*goa.Controller
}

// NewV1Controller creates a v1 controller.
func NewV1Controller(service *goa.Service) *V1Controller {
	return &V1Controller{Controller: service.NewController("V1Controller")}
}

// Random runs the random action.
func (c *V1Controller) Random(ctx *app.RandomV1Context) error {
	// V1Controller_Random: start_implement

	// Put your logic here

	return nil
	// V1Controller_Random: end_implement
}
