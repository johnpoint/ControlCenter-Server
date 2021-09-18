BIN_FILE=ControlCenter
PROJECTNAME=ControlCenter

## make: 格式化 go 代码，并编译生成二进制文件
all: check build

## build: 编译go代码生成二进制文件
build:
	@go mod tidy
	@go build -o "${BIN_FILE}" ControlCenter.go

## clean: 清理中间目标文件
clean:
	@go clean

## check: 格式化go代码
check:
	@go fmt ./
	@go vet ./

## lint: 执行代码检查
lint:
	golangci-lint run --enable-all

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'