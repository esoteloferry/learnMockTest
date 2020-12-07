package handlers

import (
	"net/http"
	"testing"

	"github.com/esoteloferry/learnMockTest/mockServices"
	"github.com/esoteloferry/learnMockTest/services/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_should_return_all_projects_with_status_code_200(t *testing.T){
  assert:=assert.New(t)
  // create an instance of our test object
  mockService:= mockServices.MockProjects{}
  // setup expectations
  mockedProjects:=[]dto.ProjectsResponse{
    {
      Id: "123",
      Name: "Test1",
    },
  }
  mockService.On("GetAllProjects").Return(mockedProjects, nil)
  // call the code we are testing
  // SETUP API
  app:= fiber.New()
  mockHandler:= Projects{
    Service: &mockService,
  }
  app.Get("/getprojects", mockHandler.GetAll)
  // SETUP REQUEST
  // Create a new http request with the route
  // from the test case
  req, _ := http.NewRequest(
    "GET",
    "/getprojects",
    nil,
  )
  // Perform the request plain with the app.
  // The -1 disables request latency.
  res, err := app.Test(req, -1)
  // verify that no error occured, that is not expected
  assert.NotNil(err, "There should not be an error")
  // Verify if the status code is as expected
  assert.Equal(200, res.StatusCode, "Status ok")
  // assert that the expectations were met
  mockService.AssertExpectations(t)
}
