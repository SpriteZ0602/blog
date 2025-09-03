# 个人博客

### 项目介绍

个人博客系统，前端使用Vue，后端使用Go来构建，集成了多种现代化特性和服务。

### 项目结构

```
├── server                              # 后端
│   ├── config/                         # 配置文件
│   │   ├── config_xxx.go               # 各种服务配置文件
│   │   └── enter.go                    # 配置入口文件
│   ├── core/                           # 核心功能
│   │   ├── config.go                   # 配置加载
│   │   ├── server.go                   # 服务器核心
│   │   └── zap.go                      # 日志处理
│   ├── global/                         # 全局变量
│   │   └── global.go                   # 全局变量定义
│   ├── initialize/                     # 初始化
│   │   └── router.go                   # 路由初始化
│   ├── model/                          # 数据模型
│   │   └── appTypes/                   # 应用类型
│   │       └── image_storage.go        # 图像存储位置判断
│   ├── utils/                          # 工具函数
│   │   └── yaml.go                     # YAML配置工具
│   ├── config.yaml                     # 配置文件
│   └── go.mod & go.sum                 # Go模块依赖
└── README.md                           # 项目说明文档
```

### 主要功能

- 用户认证 (JWT)
- 验证码支持
- 邮件服务
- 文件上传（本地/七牛云）
- QQ第三方登录
- Redis缓存
- Elasticsearch搜索
- 高德地图服务集成
- Zap日志系统

### 技术栈

- 后端：Go
- 前端：Vue
- 数据库：MySQL
- 缓存：Redis
- 搜索引擎：Elasticsearch
- 对象存储：七牛云
- 日志：Zap

### 开发环境要求

- Go 1.16+
- MySQL 8.0+
- Redis
- Elasticsearch

### 快速开始

1. 克隆项目
```bash
git clone https://github.com/SpriteZ0602/blog.git
```

2. 安装依赖
```bash
cd server
go mod tidy
```

3. 配置文件
- 复制并修改配置文件 config.yaml

4. 运行项目
```bash
go run main.go
```
