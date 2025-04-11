#!/bin/bash

# 设置变量
IMAGE_NAME="travel-guide"
IMAGE_TAG="latest"
CONTAINER_NAME="travel-guide-container"
HOST_PORT=8914
CONTAINER_PORT=8914
DATA_DIR="./data"

# 确保数据目录存在
mkdir -p $DATA_DIR

# 检查镜像是否存在
if ! docker images ${IMAGE_NAME}:${IMAGE_TAG} | grep -q ${IMAGE_NAME}; then
    echo "错误：镜像 ${IMAGE_NAME}:${IMAGE_TAG} 不存在！"
    echo "请先运行 ./build.sh 构建镜像"
    exit 1
fi

# 停止并删除已存在的容器
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "停止并删除已存在的容器..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# 运行新容器
echo "启动新容器..."
docker run -d \
    --name $CONTAINER_NAME \
    -p $HOST_PORT:$CONTAINER_PORT \
    -v $DATA_DIR:/app/data \
    --restart unless-stopped \
    ${IMAGE_NAME}:${IMAGE_TAG}

# 检查容器是否成功启动
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "容器已成功启动！"
    echo "访问 http://localhost:$HOST_PORT 查看网站"
    
    # 显示容器日志
    echo "容器日志："
    docker logs $CONTAINER_NAME
else
    echo "容器启动失败，请检查日志："
    docker logs $CONTAINER_NAME
    exit 1
fi 