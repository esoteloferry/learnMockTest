# Learning fiber + gomock!

## Only to run test
`./test`

## Folder structure
```
.
├── README.md
├── api                  *Test of the handler, it uses the mocks to isolate the handler
│   ├── handlers
│   │   ├── projects.go
│   │   └── projects_test.go
│   └── router
│       └── projects.go
├── go.mod
├── go.sum
├── mocks                 *Mock of the services
│   └── services
│       └── mockProjectsService.go
├── services
│   ├── dto               *Data transfer objects to api user
│   │   └── projects.go
│   └── projects.go
```

