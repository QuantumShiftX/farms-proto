#!/bin/bash
# genrpc.sh

# Turn on command echoing
#set -x

# 参数1: 服务名称
if [ $# -lt 1 ]; then
    echo "Usage: $0 <service-name>"
    exit 1
fi

SERVICE=$1
TARGET_DIR="../farms-rpc-server/app"
RPC_CLIENT_DIR="./proto-gen-go/${SERVICE}/v1"
# 检查目标目录是否存在，不存在则创建
mkdir -p ${TARGET_DIR}/${SERVICE}

# 检查proto文件是否存在
if [ ! -f "./proto/${SERVICE}/v1/${SERVICE}.proto" ]; then
    echo "Error: Proto file ./proto/${SERVICE}/v1/${SERVICE}.proto not found"
    exit 1
fi

# 生成rpc服务代码
goctl rpc protoc ./proto/${SERVICE}/v1/${SERVICE}.proto \
    --proto_path=. \
    --go_out=${TARGET_DIR}/${SERVICE} \
    --go-grpc_out=${TARGET_DIR}/${SERVICE} \
    --zrpc_out=${TARGET_DIR}/${SERVICE} \
    --client=true \
    --style=go_zero -m


cp -rf ${TARGET_DIR}/${SERVICE}/client/* ${RPC_CLIENT_DIR}/
# 删除源文件
rm -rf ${TARGET_DIR}/${SERVICE}/client
# 替换导入路径 - 使用变量实现动态替换
if [[ "$OSTYPE" == "darwin"* ]]; then
    # MacOS
    find ${RPC_CLIENT_DIR} -type f -name "*.go" -exec sed -i '' \
        "s|farms-rpc-server/app/${SERVICE}/pb/v1|github.com/QuantumShiftX/farms-proto/proto-gen-go/${SERVICE}/v1|g" \
        {} \;
else
    # Linux
    find ${RPC_CLIENT_DIR} -type f -name "*.go" -exec sed -i \
        "s|farms-rpc-server/app/${SERVICE}/pb/v1|github.com/QuantumShiftX/farms-proto/proto-gen-go/${SERVICE}/v1|g" \
        {} \;
fi
