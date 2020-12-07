# Learning fiber + testify/mock !

## Only to run test
`go test ./... -v`

## Folder structure
.
├── api
│   ├── handlers  *Test of the handler, it uses the mockService to isolate the handler
│   │   ├── projects.go
│   │   └── projects_test.go
│   └── router
│       └── projects.go
├── go.mod
├── go.sum
├── mockServices  *Mock of the services
│   └── projects.go
└── services
    ├── dto       *Data transfer objects to api user
    │   └── projects.go
    └── projects.go


