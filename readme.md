# goframe
Spin up small go projects.

This repo contains a boiler code to start the development of a golang backend project.

- Set up deployment
- Set up testing
- Sets up http2 SSE
- Sets up sqlite as the db
- documentation

# Dependencies
- Sqlite

## Usage
- Run `go install github.com/eltneg/goframe`
- Run `goframe newproject $projectName`
- Run `goframe newcontroller /path/to/$modelName`
- Run `cd $projectName`
- Run `go run ./server/server.go`

