# docker-golang

This is a basic project to run Golang App inside a local docker with an watcher to restart app when some file in the project is updated.

### Resource

- wtc: https://github.com/rafaelsq/wtc (watcher)
- golangci-lint: https://golangci-lint.run/ (lint)

### Run

- create `.env` file based on `.env.example`. You should use `PORT=8083` because `Dockerfile-dev` expose this port. If you changed this port, you'd also change Dockerfile.
- `docker-compose up`: create a docker container and run the app inside it.

### Extra

- `docker build -t docker-golang-app -f Dockerfile-prod .`: create a minimal docker image with app binary to execute app, without watcher or lint.
