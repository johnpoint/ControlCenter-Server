rm proto/serverInfo/*.go
protoc --go_out=../ proto/serverInfo/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/serverInfo/*.proto

rm proto/tcpProto/*.go
protoc --go_out=../ proto/tcpProto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/tcpProto/*.proto
