#grpc-client

Run protoc command:
```
protoc -I api/proto --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false api/proto/news.proto
```

Run rest server:
`go run cmd/main.go`

GRPC server for this client is [here]()

![img](vectorpaint.png)