rm proto/controlProto/*.go
protoc --go_out=../ proto/controlProto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/controlProto/*.proto

rm proto/mqProto/*.go
protoc --go_out=../ proto/mqProto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/mqProto/*.proto

rm proto/tcpProto/*.go
protoc --go_out=../ proto/tcpProto/*.proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ proto/tcpProto/*.proto
