all: build-go build-dockerfile docker-compose

clean-all: stop-docker-compose rm-build-go rm-build-dockerfile

clean-docker: stop-docker-compose rm-build-dockerfile

.PHONY: all

### build
build-go:
	cd cmd/init && go run *.go
	go build -o api-service cmd/api-service/*.go
	go build -o db-service cmd/db-service/*.go

build-dockerfile:
	docker build -t god-jay-mysql --target god-jay-mysql .
	docker build -t god-jay-server --target god-jay-server .
	docker build -t god-jay-client --target god-jay-client .

docker-compose:
	docker compose up -d

### destroy
stop-docker-compose:
	docker compose down

rm-build-go:
	rm -rf api-service db-service ./3rdparty ./cmd/init/deps ./pb/api.* ./pb/db.*

rm-build-dockerfile:
	docker rmi god-jay-mysql || true
	docker rmi god-jay-client || true
	docker rmi god-jay-server || true

