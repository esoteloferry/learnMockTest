package dto


type ProjectsResponse struct{
  Id string `json:"id"`
  Name string `json:"name"`
}

type AppError struct{
  Code int `json:",omitempty"`
  Message string `json:"message"`
}
func NewUnexpectedError()*AppError{
  return &AppError{Code:500}
}
