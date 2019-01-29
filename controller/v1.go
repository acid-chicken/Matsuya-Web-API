package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	mgo "gopkg.in/mgo.v2"
)

// V1Controller implements the v1 resource.
type V1Controller struct {
	*goa.Controller
	db *mgo.Session
}

// NewV1Controller creates a v1 controller.
func NewV1Controller(service *goa.Service, db *mgo.Session) *V1Controller {
	return &V1Controller{
		Controller: service.NewController("V1Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V1Controller) Random(ctx *app.RandomV1Context) error {
	// V1Controller_Random: start_implement

	// Put your logic here

	return nil
	// V1Controller_Random: end_implement
}
