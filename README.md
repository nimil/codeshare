# 旅游攻略分享平台

一个使用 Go + Gin 构建的旅游攻略分享平台，支持展示多个城市的旅游攻略。基于高德地图 MCP Server 和 AI 模型自动生成精美的旅游攻略。

## 功能特点

- 响应式设计，支持多设备访问
- 精美的卡片式布局
- 支持多个攻略页面
- 集成高德地图 MCP 服务
- AI 自动生成攻略内容
- Docker 容器化部署

## 效果展示

![首页预览](docs/images/home.png)
![西安攻略](docs/images/xian.png)
![山东攻略](docs/images/shandong.png)

## 快速开始

### 获取代码

```bash
# 克隆仓库
git clone git@github.com:nimil/codeshare.git
cd codeshare
```

### 配置高德地图 MCP

1. 注册高德开放平台账号并创建应用
2. 获取 Key 和 Secret
3. 在 Cursor 中配置 MCP 工具：
```json
{
  "mcp_amap": {
    "key": "your_key",
    "secret": "your_secret"
  }
}
```

### 生成旅游攻略

1. 使用高德 MCP 工具获取地理信息：
```python
# 示例：获取景点信息
response = mcp_amap_text_search(keywords="西安城墙", city="西安")
```

2. 生成攻略提示词示例：
```text
为{城市名}生成一个{天数}日游攻略，包含以下内容：
- 每日行程安排
- 特色美食推荐
- 交通建议
- 住宿推荐
- 注意事项
```

3. 使用 DeepSeek-V3 生成 HTML：
```text
将以下旅游攻略转换为精美的响应式 HTML 页面：
- 使用现代化的 UI 设计
- 添加卡片式布局
- 包含动态效果
- 适配移动端
```

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
├── docs/             # 文档和图片
│   └── images/       # 效果图
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
- 高德地图 MCP Server
- DeepSeek-V3
- Docker

## API 文档

### 高德地图 MCP 接口

本项目使用了以下高德地图 MCP 接口：

- 地理编码：将地址转换为经纬度
- 关键词搜索：搜索景点、餐厅等 POI 信息
- 路径规划：提供驾车、步行、公交等出行方案
- 天气查询：获取目的地天气信息
- 周边搜索：查找景点周边的设施

详细 API 文档请参考：[高德地图 MCP Server 文档](https://lbs.amap.com/api/mcp-server/summary)

## 端口说明

服务默认运行在 8914 端口。

## 许可证

MIT License 