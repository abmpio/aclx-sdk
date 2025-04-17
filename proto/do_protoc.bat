@echo off
setlocal

:: 设置 PATH 环境变量
set PATH=%PATH%;%USERPROFILE%\go\bin

:: 运行 protoc 命令
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./aclx.proto

endlocal