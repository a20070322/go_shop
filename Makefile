SHELL=/usr/bin/env bash
uname=${shell uname}
appName=go_shop
build_path= ./dist/
linux_app_run_path = /app/go_shop_linux/
linux_app_name = go-shop
linux_scp_port = 6079
linux_scp_user = root
linux_scp_host = 49.232.168.44

all: build

build: linux mac windows
deploy_44: linux deploy

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

ent_m:
	@go run main.go ent_m

deploy_init:
	@echo "部署初始化文件...."
	@scp -P ${linux_scp_port} -r ./config.yaml ${linux_scp_user}@${linux_scp_host}:${linux_app_run_path}
	@scp -P ${linux_scp_port} -r ./cert ${linux_scp_user}@${linux_scp_host}:${linux_app_run_path}
deploy:
	@echo "部署项目...."
	@ssh -p ${linux_scp_port} ${linux_scp_user}@${linux_scp_host} "systemctl stop ${linux_app_name}"
	@scp -P ${linux_scp_port} -r ./dist/${appName}_linux ${linux_scp_user}@${linux_scp_host}:${linux_app_run_path}
	@ssh -p ${linux_scp_port} ${linux_scp_user}@${linux_scp_host} "systemctl start ${linux_app_name}"
	@echo "部署成功"
deploy_service:
	@echo "部署服务文件...."
	@scp -P 6079 -r ./${linux_app_name}.service ${linux_scp_user}@${linux_scp_host}:/etc/systemd/system/