#!/bin/bash

# 设置镜像名称和标签
IMAGE_NAME="travel-guide"
IMAGE_TAG="latest"

# 构建适用于 amd64 架构的 Docker 镜像
echo "开始构建 Docker 镜像..."
docker buildx build \
    --platform linux/amd64 \
    -t ${IMAGE_NAME}:${IMAGE_TAG} \
    --load \
    .

# 检查构建是否成功
if [ $? -eq 0 ]; then
    echo "Docker 镜像构建成功！"
    
    # 导出镜像为 tar 文件
    echo "正在导出镜像文件..."
    docker save ${IMAGE_NAME}:${IMAGE_TAG} -o ${IMAGE_NAME}.tar
    
    echo "镜像已保存为 ${IMAGE_NAME}.tar"
    echo "在 CentOS 上使用以下命令部署："
    echo "# 1. 加载镜像"
    echo "docker load -i ${IMAGE_NAME}.tar"
    echo
    echo "# 2. 清理同名容器（如果存在）"
    echo "docker stop ${IMAGE_NAME} 2>/dev/null"
    echo "docker rm ${IMAGE_NAME} 2>/dev/null"
    echo
    echo "# 3. 启动新容器"
    echo "docker run -d --name ${IMAGE_NAME} -p 8914:8914 ${IMAGE_NAME}:${IMAGE_TAG}"
else
    echo "Docker 镜像构建失败！"
    exit 1
fi 