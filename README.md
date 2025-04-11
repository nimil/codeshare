# 旅游攻略分享平台

一个使用 Go + Gin 构建的旅游攻略分享平台，支持展示多个城市的旅游攻略。

## 功能特点

- 响应式设计，支持多设备访问
- 精美的卡片式布局
- 支持多个攻略页面
- Docker 容器化部署

## 快速开始

### 构建 Docker 镜像

```bash
# 构建 amd64 架构的 Docker 镜像
./build.sh
```

### 部署运行

```bash
# 在目标服务器上运行
./deploy.sh
```

或者手动执行以下步骤：

1. 加载 Docker 镜像：
```bash
docker load -i travel-guide.tar
```

2. 清理同名容器（如果存在）：
```bash
docker stop travel-guide
docker rm travel-guide
```

3. 启动新容器：
```bash
docker run -d --name travel-guide -p 8914:8914 travel-guide:latest
```

## 目录结构

```
.
├── Dockerfile          # Docker 构建文件
├── README.md          # 项目说明文档
├── build.sh           # 构建脚本
├── deploy.sh          # 部署脚本
├── go.mod            # Go 模块文件
├── go.sum            # Go 依赖版本锁定文件
├── main.go           # 主程序入口
├── static/           # 静态资源目录
└── templates/        # HTML 模板目录
    ├── 404.html
    ├── index.html
    ├── shandonglaodong.html
    ├── view.html
    └── xianthree.html
```

## 技术栈

- Go 1.20
- Gin Web 框架
- HTML5 + CSS3
- Docker

## 端口说明

服务默认运行在 8914 端口。

## 许可证

MIT License 