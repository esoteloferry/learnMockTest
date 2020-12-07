package handlers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/esoteloferry/learnMockTest/services"
)

type Projects struct{
  Service services.IProjects
}

func (ph *Projects) GetAll(c *fiber.Ctx) error{
  projects, err:= ph.Service.GetAllProjects()
  if err!= nil{
    return c.Status(err.Code).JSON(err)
  }
  c.JSON(projects)
  return nil
}
