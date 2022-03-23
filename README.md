# docker-golang

This is a basic project to run Golang App inside a local docker with an watcher to restart app when some file in the project is updated.

### Resource

- wtc: https://github.com/rafaelsq/wtc (watcher)
- golangci-lint: https://golangci-lint.run/ (lint)

### Run

- `docker-compose up`: create a docker and run the app inside the docker. the app listens port 8083 and the docker expose port 8083 too.

### Extra

- `docker build -t docker-golang-app -f Dockerfile-prod .`: create a minimal docker image with app binary to execute app, without watcher or lint.
