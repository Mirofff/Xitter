Для генерации grpc файлов, используется protoc команда:

```sh
grpc_out=src/interfaces/proto grpc_opt="paths=source_relative" protoc -I ../../third_party/googleapis -I $grpc_out --go-grpc_out
 $grpc_out --go_out=$grpc_out --go_opt=$grpc_opt --go-grpc_opt=$grpc_opt --grpc-gateway_out $grpc_out --grpc-gateway_opt=$grpc_opt --openapiv2_out $grpc_out --grpc-gateway_opt generate_unbound_methods=true
 src/interfaces/proto/*.proto
```
