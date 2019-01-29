package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	mgo "gopkg.in/mgo.v2"
)

// V3Controller implements the v3 resource.
type V3Controller struct {
	*goa.Controller
	db *mgo.Database
}

// NewV3Controller creates a v3 controller.
func NewV3Controller(service *goa.Service, db *mgo.Database) *V3Controller {
	return &V3Controller{
		Controller: service.NewController("V3Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V3Controller) Random(ctx *app.RandomV3Context) error {
	// V3Controller_Random: start_implement

	// Put your logic here

	res := &app.MeMakotiaMatsuyaV3{}
	return ctx.OK(res)
	// V3Controller_Random: end_implement
}
