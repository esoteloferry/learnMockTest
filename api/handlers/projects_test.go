package handlers

import (
	"net/http"
	"testing"

	"github.com/esoteloferry/learnMockTest/mocks/services"
	"github.com/esoteloferry/learnMockTest/services/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

var app *fiber.App
var ph Projects
var mockService *services.MockIProjects

func setup(t *testing.T)func(){
  ctrl:= gomock.NewController(t)
  mockService= services.NewMockIProjects(ctrl)

  ph = Projects{
    Service: mockService,
  }
  app= fiber.New()
  return func(){
    app= nil//Reset app
    defer ctrl.Finish()
  }
}

func Test_should_return_all_projects_with_status_code_200(t *testing.T){
  // Arrange 
  teardown:=setup(t)
  defer teardown()
  // Expectations
  MockedProjects:=[]dto.ProjectsResponse{
      {
        Id: "123",
        Name: "Test1",
      },
    }
  mockService.EXPECT().GetAllProjects().Return(MockedProjects, nil)
  // call the code we are testing
  app.Get("/getprojects", ph.GetAll)
  req,_:=http.NewRequest(http.MethodGet, "/getprojects", nil)
  res, err:= app.Test(req, -1)
  // Assert
  if res.StatusCode!= http.StatusOK{
    t.Error("Failed while testing status code")
  }
  if err!=nil{
    t.Error("Some error happened")
  }
}
func Test_should_return_status_code_500_with_error_message(t *testing.T){
  // Arrange 
  teardown:=setup(t)
  defer teardown()
  // Expectations
  mockService.EXPECT().GetAllProjects().Return(nil,dto.NewUnexpectedError())
  // call the code we are testing
  app.Get("/getprojects", ph.GetAll)
  req,_:=http.NewRequest(http.MethodGet, "/getprojects", nil)
  res, err:= app.Test(req, -1)
  // Assert
  if res.StatusCode!= http.StatusInternalServerError{
    t.Error("Failed while testing status code")
  }
  if err!=nil{
    t.Error("Some error happened")
  }
}
