package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	mgo "gopkg.in/mgo.v2"
)

// V2Controller implements the v2 resource.
type V2Controller struct {
	*goa.Controller
	db *mgo.Session
}

// NewV2Controller creates a v2 controller.
func NewV2Controller(service *goa.Service, db *mgo.Session) *V2Controller {
	return &V2Controller{
		Controller: service.NewController("V2Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V2Controller) Random(ctx *app.RandomV2Context) error {
	// V2Controller_Random: start_implement

	// Put your logic here

	res := &app.MeMakotiaMatsuyaV2{}
	return ctx.OK(res)
	// V2Controller_Random: end_implement
}
