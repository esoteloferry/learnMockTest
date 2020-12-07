package mockServices

import (
	"github.com/stretchr/testify/mock"
  "github.com/esoteloferry/learnMockTest/services/dto"
)


type MockProjects struct{
  mock.Mock
}

//Implement methods to comply interface of IProjects
func(p *MockProjects) GetAllProjects()([]dto.ProjectsResponse, *dto.AppError){
  args:= p.Called()
  return args.Get(0).([]dto.ProjectsResponse), args.Get(1).(*dto.AppError)
}
