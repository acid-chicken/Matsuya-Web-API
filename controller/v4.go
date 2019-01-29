package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	mgo "gopkg.in/mgo.v2"
)

// V4Controller implements the v4 resource.
type V4Controller struct {
	*goa.Controller
	db *mgo.Session
}

// NewV4Controller creates a v4 controller.
func NewV4Controller(service *goa.Service, db *mgo.Session) *V4Controller {
	return &V4Controller{
		Controller: service.NewController("V4Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V4Controller) Random(ctx *app.RandomV4Context) error {
	// V4Controller_Random: start_implement

	// Put your logic here

	res := &app.MeMakotiaMatsuyaRandom{}
	return ctx.OK(res)
	// V4Controller_Random: end_implement
}

// Search runs the search action.
func (c *V4Controller) Search(ctx *app.SearchV4Context) error {
	// V4Controller_Search: start_implement

	// Put your logic here

	return nil
	// V4Controller_Search: end_implement
}
