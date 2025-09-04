# 个人博客

### 项目介绍

个人博客系统，采用前后端分离架构，后端使用 Go 语言开发，集成了多种现代化特性和服务，包括 Redis 缓存、ElasticSearch 搜索引擎、Gorm ORM框架、Zap 日志系统等。

### 系统架构

1. 数据存储层
   - MySQL：核心数据存储
   - Redis：高速缓存服务
   - ElasticSearch：全文检索引擎
   
2. 应用服务层
   - Web服务：基于 Go 的 HTTP 服务
   - 定时任务：Cron 任务调度系统
   - 日志系统：Zap 日志框架

3. 基础设施
   - 配置中心：YAML 配置管理
   - ORM：Gorm 数据库操作
   - 缓存：Redis 数据缓存

### 项目结构

```
server/                              # 后端目录
├── config/                          # 配置文件目录
│   ├── config_xxx.go                # 读取各种配置
│   └── enter.go                     # 配置入口
├── core/                            # 核心功能
│   ├── config.go                    # 配置加载
│   ├── server.go                    # 服务器核心
│   └── zap.go                       # 日志处理
├── global/                          # 全局对象
│   └── global.go                    # 全局实例定义
├── initialize/                      # 初始化模块
│   ├── cron.go                      # 定时任务初始化
│   ├── es.go                        # ES 初始化
│   ├── gorm.go                      # 数据库初始化
│   ├── others.go                    # 其他初始化
│   ├── redis.go                     # Redis 初始化
│   └── router.go                    # 路由初始化
├── model/                           # 数据模型
│   └── appTypes/                    # 应用类型
│       └── image_storage.go         # 存储类型
├── task/                            # 定时任务
│   └── enter.go                     # 任务入口
├── utils/                           # 工具函数
│   ├── parse.go                     # 解析工具
│   └── yaml.go                      # YAML工具
├── log/                             # 日志目录
├── config.yaml                      # 配置文件
└── go.mod & go.sum                  # 依赖管理
```
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
