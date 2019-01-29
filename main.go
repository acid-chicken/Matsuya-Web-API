//go:generate goagen bootstrap -d github.com/makotia/Matsuya-Web-API/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/makotia/Matsuya-Web-API/app"
	"github.com/makotia/Matsuya-Web-API/controller"
)

func main() {
	// Create service
	service := goa.New("Matsuya-Web-API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "v1" controller
	c := controller.NewV1Controller(service)
	app.MountV1Controller(service, c)
	// Mount "v2" controller
	c2 := controller.NewV2Controller(service)
	app.MountV2Controller(service, c2)
	// Mount "v3" controller
	c3 := controller.NewV3Controller(service)
	app.MountV3Controller(service, c3)
	// Mount "v4" controller
	c4 := controller.NewV4Controller(service)
	app.MountV4Controller(service, c4)

	// Start service
	if err := service.ListenAndServeTLS(":8080", "cert.pem", "key.pem"); err != nil {
		service.LogError("startup", "err", err)
	}

}
