## github.com/god-jay/gools/discovery 包的使用示例

### 用法

你需要先安装 docker, docker compose

1. 进入 `my-micro-service` 目录
2. 在终端里面输入 `make` 命令（使用的：[Makefile](Makefile)）
   - 这会 build go, build docker images（使用的：[Dockerfile](Dockerfile)），
     然后 run docker 容器（使用的：[docker-compose.yml](docker-compose.yml)）
   - api docker 容器会为 api 服务暴露一个8080端口，如果你本地8080端口被占用了，
     你可以任意修改这个端口（[`docker-compose.yml`](docker-compose.yml) 文件里的第一个8080数字）
3. 发送 http 请求来调用 api 接口：
   - `curl -X POST --location "http://localhost:8080/user" -H "Content-Type: application/json" -d "{\"id\": 1}"`
   - 或者你可以直接在终端里面运行 `make run-k6` 命令
4. 你可以 stop/start 任意一个 `god-jay-micro-db-service` docker 容器，
   它会被自动 删除/添加 到后端服务中，通过 discovery 包的服务发现功能

## 注意

在这个示例里，我删除了所有 proto 文件生成的 go 文件，并且保留了逻辑代码，然后需要用 `make` 命令通过 `cmd/init/main.go` 重新生成这些文件

不过在你的实际开发中，你应该先写 proto 文件，然后通过命令生成需要的 go 文件，最后再写你的逻辑代码