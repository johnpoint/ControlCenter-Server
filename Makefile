BIN_FILE=ControlCenter

all: check build

build:
	@go mod tidy
	@go build -o "${BIN_FILE}" ControlCenter.go

clean:
	@go clean

check:
	@go fmt ./
	@go vet ./

lint:
	golangci-lint run --enable-all

help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make check 格式化go代码"
	@echo "make lint 执行代码检查"