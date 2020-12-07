package router

import (
  "github.com/esoteloferry/learnMockTest/services"
  "github.com/esoteloferry/learnMockTest/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Add(app *fiber.App){
  //WIRING
  ph := handlers.Projects{
      Service: services.NewProjects(),
    }
  //ROUTES
  app.Get("getprojects", ph.GetAll)
}
