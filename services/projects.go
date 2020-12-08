package services

import (
  "github.com/esoteloferry/learnMockTest/services/dto"
)
//go:generate mockgen -destination=../mocks/services/mockProjectsService.go -package=services github.com/esoteloferry/learnMockTest/services IProjects
type IProjects interface{
  GetAllProjects() ([]dto.ProjectsResponse, *dto.AppError)
}
type Projects struct{
}
//Implement the methods of interface
func (p Projects) GetAllProjects()(r []dto.ProjectsResponse,err *dto.AppError){
  //Here the code interacts with the real database 
  //but for now we are making something fake to not put many files in these example
  r= []dto.ProjectsResponse{
    {
      Id: "111111",
      Name: "Test1",
    },
  }
  return r, nil
}

//Initialize
func NewProjects ()Projects{
  return Projects{}
}
