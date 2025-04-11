#!/bin/bash

IMAGE_NAME="travel-guide"
IMAGE_TAG="latest"

# 1. 加载镜像
echo "加载 Docker 镜像..."
docker load -i ${IMAGE_NAME}.tar

# 2. 清理同名容器
echo "清理旧容器..."
docker stop ${IMAGE_NAME} 2>/dev/null
docker rm ${IMAGE_NAME} 2>/dev/null

# 3. 启动新容器
echo "启动新容器..."
docker run -d --name ${IMAGE_NAME} -p 8914:8914 ${IMAGE_NAME}:${IMAGE_TAG}

echo "部署完成！服务运行在 http://localhost:8914" 