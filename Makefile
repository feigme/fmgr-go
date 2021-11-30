# Makefile 介绍 https://seisman.github.io/how-to-write-makefile/index.html

# 使用#> 规则自动获取命令介绍
help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^#>//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	
#> make build: 编译Go代码, 生成二进制文件
build:
	@go build ./
  
#> make run: 直接运行Go代码
run:
	@go run ./
  
#> make test: 运行测试用例
test:
	@go test ./...	  

#> make build-snapshot: 构建snapshot各个环境的二进制文件
build-snapshot:
	@goreleaser --snapshot --skip-publish --rm-dist	

.PHONY: build run help test
