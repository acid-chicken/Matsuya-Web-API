//go:generate goagen bootstrap -d github.com/makotia/Matsuya-Web-API/design

package main

import (
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/makotia/Matsuya-Web-API/app"
	"github.com/makotia/Matsuya-Web-API/controller"
	"github.com/makotia/Matsuya-Web-API/utils"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// Create service
	service := goa.New("Matsuya-Web-API")

	utils.LoadEnv()
	mgoSession, err := mgo.Dial(utils.GetMongoConnectionURL())
	if err != nil {
		log.Fatal(err)
	}
	defer mgoSession.Close()

	db := mgoSession.DB(utils.GetDBName())

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "v1" controller
	c := controller.NewV1Controller(service, db)
	app.MountV1Controller(service, c)
	// Mount "v2" controller
	c2 := controller.NewV2Controller(service, db)
	app.MountV2Controller(service, c2)
	// Mount "v3" controller
	c3 := controller.NewV3Controller(service, db)
	app.MountV3Controller(service, c3)
	// Mount "v4" controller
	c4 := controller.NewV4Controller(service, db)
	app.MountV4Controller(service, c4)

	// Start service
	if err := service.ListenAndServeTLS(":8080", "cert.pem", "key.pem"); err != nil {
		service.LogError("startup", "err", err)
	}

}
