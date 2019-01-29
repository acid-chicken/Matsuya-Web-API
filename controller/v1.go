package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	"github.com/makotia/Matsuya-Web-API/repository"
	mgo "gopkg.in/mgo.v2"
)

// V1Controller implements the v1 resource.
type V1Controller struct {
	*goa.Controller
	db *mgo.Database
}

// NewV1Controller creates a v1 controller.
func NewV1Controller(service *goa.Service, db *mgo.Database) *V1Controller {
	return &V1Controller{
		Controller: service.NewController("V1Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V1Controller) Random(ctx *app.RandomV1Context) error {
	// V1Controller_Random: start_implement

	// Put your logic here
	r := repository.NewMenuRepository(c.db)
	menu, err := r.FindRandom()
	if err != nil {
		return ctx.InternalServerError()
	}

	return ctx.OK([]string{menu.Name})
	// V1Controller_Random: end_implement
}
