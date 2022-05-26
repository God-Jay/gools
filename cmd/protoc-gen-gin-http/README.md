A protoc plugin to generate gin HTTP handlers from proto files.
This will generate a `.gin.http.pb.go` file for each `.proto` file.

By using this plugin, once you wrote your `.proto` file,
you can generate its gin HTTP handlers,
which contains the routes and the request and response types defined in the `.proto` file.
You will only need to focus on handling the request and response logic.

## example

- [gin-protoc-project](../../_examples/gin-protoc-project)

## This plugin is developed based on the go-kratos project:

[https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http](https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http)

Thanks for your jobs.