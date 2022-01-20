# Makefile 介绍 https://seisman.github.io/how-to-write-makefile/index.html

# 使用#> 规则自动获取命令介绍
.PHONY: help
help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^#>//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	
#> make build: 编译Go代码, 生成二进制文件
.PHONY: build
build:
	@go build ./
  
#> make run: 直接运行Go代码
.PHONY: run
run:
	@go run ./
  
#> make test: 运行测试用例
.PHONY: test
test:
	@go test -v ./...	  

#> make build-snapshot: 构建snapshot各个环境的二进制文件
.PHONY: build-snapshot
build-snapshot:
	@goreleaser --snapshot --skip-publish --rm-dist	
