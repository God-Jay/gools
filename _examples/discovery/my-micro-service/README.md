## This is an example of using the github.com/god-jay/gools/discovery package.

### Usage

You must install docker, docker compose first.

1. go to the `my-micro-service` directory
2. run `make` command in your terminal (use: [Makefile](Makefile))
    - this command will build go, build docker images (use: [Dockerfile](Dockerfile)),
      and run docker container (use: [docker-compose.yml](docker-compose.yml))
    - the api docker container will expose a 8080 port for the api server, fill free to change the port (the
      first `8080` in the [`docker-compose.yml`](docker-compose.yml) file) if it is used
3. send http request to make the api request:
   - `curl -X POST --location "http://localhost:8080/user" -H "Content-Type: application/json" -d "{\"id\": 1}"`
   - or you can simply run `make run-k6` command in your terminal
4. feel free to stop/start one of the `god-jay-micro-db-service` docker container,
   it can be removed/added to the backend server automatically by the discovery service

## Notice

In this example, I deleted all the generated proto go files and kept the logic codes, then generate them with `make` by
the `cmd/init/main.go` code.

But in your real develop, you should write proto files first, then generate the go files,
finally write your logic codes.