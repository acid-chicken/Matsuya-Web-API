package controller

import (
	"math/rand"
	"time"

	"github.com/goadesign/goa"
	"github.com/makotia/Matsuya-Web-API/app"
	"github.com/makotia/Matsuya-Web-API/repository"
	mgo "gopkg.in/mgo.v2"
)

// V4Controller implements the v4 resource.
type V4Controller struct {
	*goa.Controller
	db *mgo.Database
}

// NewV4Controller creates a v4 controller.
func NewV4Controller(service *goa.Service, db *mgo.Database) *V4Controller {
	return &V4Controller{
		Controller: service.NewController("V4Controller"),
		db:         db,
	}
}

// Random runs the random action.
func (c *V4Controller) Random(ctx *app.RandomV4Context) error {
	// V4Controller_Random: start_implement

	// Put your logic here
	r := repository.NewMenuRepository(c.db)
	menu, err := r.FindRandom()
	if err != nil {
		return ctx.InternalServerError()
	}
	if ctx.Type != nil && ctx.Name != nil {
		typeFetched, err := r.FindByType(*ctx.Type)
		if err != nil {
			return ctx.InternalServerError()
		}
		nameFetched, err := r.FindByName(*ctx.Name)
		if err != nil {
			return ctx.InternalServerError()
		}
		merged := append(typeFetched, nameFetched)
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(merged))
		res := &app.MeMakotiaMatsuyaV4{
			Calorie:        merged[i].Calorie,
			Carbohydrate:   merged[i].Carbohydrate,
			Description:    merged[i].Description,
			ImageURL:       merged[i].ImageURL,
			Lipid:          merged[i].Lipid,
			Name:           merged[i].Name,
			Price:          merged[i].Price,
			Protein:        merged[i].Protein,
			SaltEquivalent: merged[i].SaltEquivalent,
			Sodium:         merged[i].Sodium,
			Type:           merged[i].Type,
		}
		return ctx.OK(res)
	}

	if ctx.Type != nil {
		typeFetched, err := r.FindByType(*ctx.Type)
		if err != nil {
			return ctx.InternalServerError()
		}
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(typeFetched))
		res := &app.MeMakotiaMatsuyaV4{
			Calorie:        typeFetched[i].Calorie,
			Carbohydrate:   typeFetched[i].Carbohydrate,
			Description:    typeFetched[i].Description,
			ImageURL:       typeFetched[i].ImageURL,
			Lipid:          typeFetched[i].Lipid,
			Name:           typeFetched[i].Name,
			Price:          typeFetched[i].Price,
			Protein:        typeFetched[i].Protein,
			SaltEquivalent: typeFetched[i].SaltEquivalent,
			Sodium:         typeFetched[i].Sodium,
			Type:           typeFetched[i].Type,
		}
		return ctx.OK(res)
	}

	if ctx.Name != nil {
		nameFetched, err := r.FindByName(*ctx.Name)
		if err != nil {
			return ctx.InternalServerError()
		}
		res := &app.MeMakotiaMatsuyaV4{
			Calorie:        nameFetched.Calorie,
			Carbohydrate:   nameFetched.Carbohydrate,
			Description:    nameFetched.Description,
			ImageURL:       nameFetched.ImageURL,
			Lipid:          nameFetched.Lipid,
			Name:           nameFetched.Name,
			Price:          nameFetched.Price,
			Protein:        nameFetched.Protein,
			SaltEquivalent: nameFetched.SaltEquivalent,
			Sodium:         nameFetched.Sodium,
			Type:           nameFetched.Type,
		}
		return ctx.OK(res)
	}

	res := &app.MeMakotiaMatsuyaV4{
		Calorie:        menu.Calorie,
		Carbohydrate:   menu.Carbohydrate,
		Description:    menu.Description,
		ImageURL:       menu.ImageURL,
		Lipid:          menu.Lipid,
		Name:           menu.Name,
		Price:          menu.Price,
		Protein:        menu.Protein,
		SaltEquivalent: menu.SaltEquivalent,
		Sodium:         menu.Sodium,
		Type:           menu.Type,
	}
	return ctx.OK(res)
	// V4Controller_Random: end_implement
}

// Search runs the search action.
func (c *V4Controller) Search(ctx *app.SearchV4Context) error {
	// V4Controller_Search: start_implement

	// Put your logic here
	r := repository.NewMenuRepository(c.db)
	if ctx.Type != nil && ctx.Name != nil {
		typeFetched, err := r.FindByType(*ctx.Type)
		if err != nil {
			return ctx.InternalServerError()
		}
		nameFetched, err := r.FindByName(*ctx.Name)
		if err != nil {
			return ctx.InternalServerError()
		}
		merged := append(typeFetched, nameFetched)
		res := make([]*app.MeMakotiaMatsuyaV4, len(merged))
		for i, item := range merged {
			v4 := &app.MeMakotiaMatsuyaV4{
				Calorie:        item.Calorie,
				Carbohydrate:   item.Carbohydrate,
				Description:    item.Description,
				ImageURL:       item.ImageURL,
				Lipid:          item.Lipid,
				Name:           item.Name,
				Price:          item.Price,
				Protein:        item.Protein,
				SaltEquivalent: item.SaltEquivalent,
				Sodium:         item.Sodium,
				Type:           item.Type,
			}
			res[i] = v4
		}
		return ctx.OK(res)
	}

	if ctx.Type != nil {
		typeFetched, err := r.FindByType(*ctx.Type)
		if err != nil {
			return ctx.InternalServerError()
		}
		res := make([]*app.MeMakotiaMatsuyaV4, len(typeFetched))
		for i, item := range typeFetched {
			v4 := &app.MeMakotiaMatsuyaV4{
				Calorie:        item.Calorie,
				Carbohydrate:   item.Carbohydrate,
				Description:    item.Description,
				ImageURL:       item.ImageURL,
				Lipid:          item.Lipid,
				Name:           item.Name,
				Price:          item.Price,
				Protein:        item.Protein,
				SaltEquivalent: item.SaltEquivalent,
				Sodium:         item.Sodium,
				Type:           item.Type,
			}
			res[i] = v4
		}
		return ctx.OK(res)
	}

	if ctx.Name != nil {
		nameFetched, err := r.FindByName(*ctx.Name)
		if err != nil {
			return ctx.InternalServerError()
		}
		res := &app.MeMakotiaMatsuyaV4{
			Calorie:        nameFetched.Calorie,
			Carbohydrate:   nameFetched.Carbohydrate,
			Description:    nameFetched.Description,
			ImageURL:       nameFetched.ImageURL,
			Lipid:          nameFetched.Lipid,
			Name:           nameFetched.Name,
			Price:          nameFetched.Price,
			Protein:        nameFetched.Protein,
			SaltEquivalent: nameFetched.SaltEquivalent,
			Sodium:         nameFetched.Sodium,
			Type:           nameFetched.Type,
		}
		return ctx.OK([]*app.MeMakotiaMatsuyaV4{res})
	}

	menu, err := r.GetAllMenu()
	if err != nil {
		return ctx.InternalServerError()
	}

	res := make([]*app.MeMakotiaMatsuyaV4, len(menu))
	for i, item := range menu {
		v4 := &app.MeMakotiaMatsuyaV4{
			Calorie:        item.Calorie,
			Carbohydrate:   item.Carbohydrate,
			Description:    item.Description,
			ImageURL:       item.ImageURL,
			Lipid:          item.Lipid,
			Name:           item.Name,
			Price:          item.Price,
			Protein:        item.Protein,
			SaltEquivalent: item.SaltEquivalent,
			Sodium:         item.Sodium,
			Type:           item.Type,
		}
		res[i] = v4
	}
	return ctx.OK(res)
	// V4Controller_Search: end_implement
}
