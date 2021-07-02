SHELL=/usr/bin/env bash
uname=${shell uname}
appName=go_shop
build_path= ./dist/

all: build

build: linux mac windows

linux:
	@echo 编译linux版本文件
	@env CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o ${build_path}${appName}_linux

mac:
	@echo 编译mac版本文件
	@if [ ${uname}="Darwin" ]; then  go build -o ${build_path}${appName}_mac; else env CGO_ENABLED=0 GOOS=darwin GOARCH=amd6 go build -o ${build_path}${appName}_mac; fi
windows:
	@echo 编译windows版本文件
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${build_path}${appName}_windows

clean:
	@rm -rf ${build_path}
	@echo 清除成功

ent_g:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema