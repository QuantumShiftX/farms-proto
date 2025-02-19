@echo off
REM genrpc.bat

REM Check if service name parameter is provided
if "%~1"=="" (
    echo Usage: %0 ^<service-name^>
    exit /b 1
)

set SERVICE=%~1
set TARGET_DIR=..\farms-rpc-server\app

REM Create target directory if it doesn't exist
if not exist "%TARGET_DIR%\%SERVICE%" mkdir "%TARGET_DIR%\%SERVICE%"

REM Check if proto file exists
if not exist ".\proto\%SERVICE%\v1\%SERVICE%.proto" (
    echo Error: Proto file .\proto\%SERVICE%\v1\%SERVICE%.proto not found
    exit /b 1
)

REM Generate RPC service code
goctl rpc protoc .\proto\%SERVICE%\v1\%SERVICE%.proto ^
    --proto_path=. ^
    --go_out=%TARGET_DIR%\%SERVICE% ^
    --go-grpc_out=%TARGET_DIR%\%SERVICE% ^
    --zrpc_out=%TARGET_DIR%\%SERVICE% ^
    --client=true ^
    --style=go_zero -m