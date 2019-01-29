package controller

import (
	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	"github.com/makotia/Matsuya-Web-API/repository"
	mgo "gopkg.in/mgo.v2"
)

// V2Controller implements the v2 resource.
type V2Controller struct {
	*goa.Controller
	db *mgo.Database
}

// NewV2Controller creates a v2 controller.
func NewV2Controller(service *goa.Service, db *mgo.Database) *V2Controller {
	return &V2Controller{
		Controller: service.NewController("V2Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V2Controller) Random(ctx *app.RandomV2Context) error {
	// V2Controller_Random: start_implement

	// Put your logic here
	r := repository.NewMenuRepository(c.db)
	menu, err := r.FindRandom()
	if err != nil {
		return ctx.InternalServerError()
	}

	res := &app.MeMakotiaMatsuyaV2{
		Menu: menu.Name,
	}
	return ctx.OK(res)
	// V2Controller_Random: end_implement
}
