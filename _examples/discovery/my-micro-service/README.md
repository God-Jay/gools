## This is an example of using the github.com/god-jay/gools/discovery package.

### Usage

1. go to the my-micro-service directory
2. type `make` command in your terminal (use: [Makefile](Makefile))
    - this command will build go, build docker images (use: [Dockerfile](Dockerfile)), run docker container (use: [docker-compose.yml](docker-compose.yml))
    - the api docker container will expose a 8080 port for the api server, fill free to change the port (the first `8080` in
      the [`docker-compose.yml`](docker-compose.yml) file) if it is used
3. send http request to make the api request:
   `curl -X POST --location "http://localhost:8080/user" -H "Content-Type: application/json" -d "{\"id\": 1}"`
   

