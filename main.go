//go:generate goagen bootstrap -d github.com/shogo82148/goa-adder/design

package main

import (
	"github.com/deadcheat/goacors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/shogo82148/goa-adder/app"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(goacors.WithConfig(service, &goacors.DefaultGoaCORSConfig))
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
