#!/bin/bash

# 设置变量
IMAGE_NAME="travel-guide"
IMAGE_TAG="latest"

# 构建新镜像
echo "构建新镜像..."
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

# 检查构建是否成功
if [ $? -eq 0 ]; then
    echo "镜像构建成功！"
    echo "镜像名称: ${IMAGE_NAME}:${IMAGE_TAG}"
    
    # 显示镜像信息
    docker images | grep ${IMAGE_NAME}
else
    echo "镜像构建失败！"
    exit 1
fi 