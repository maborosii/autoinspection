.PHONY: all build clean run check cover lint docker help

BIN_FILE=ai

# all: check build
build:
	@go build -o "${BIN_FILE}"
clean:
	@go clean
	rm --force "xx.out"
test:
	@go test ./...
check:
	@go fmt ./...
	@go vet ./...
cover:
	@go test -coverprofile xx.out
	@go tool cover -html=xx.out
run:
	./"${BIN_FILE}"
lint:
	golangci-lint run -c .golangci.yaml -v --timeout=3m0s
docker:
	@docker build -t maborosii/ai:latest .
help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make cover 检查测试覆盖率"
	@echo "make run 直接运行程序"
	@echo "make lint 执行代码检查"
	@echo "make docker 构建docker镜像"
