# gRPC client example

Run protoc command:
```
protoc -I api/proto --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false api/proto/news.proto
```

Run rest server:
`go run cmd/main.go`

GRPC server for this client is [here](https://github.com/KonstantinP85/grpc-server)

![vectorpaint](https://user-images.githubusercontent.com/74908254/145671185-1e42e3c6-940f-419d-9058-81a5e069ef0e.png)
