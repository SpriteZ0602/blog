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
├── api/                             # API 控制器
│   ├── base.go                      # 基础API控制器
│   └── enter.go                     # API入口
├── config/                          # 配置文件目录
│   ├── config_captcha.go            # 验证码配置
│   ├── config_email.go              # 邮件配置
│   ├── config_es.go                 # Elasticsearch配置
│   ├── config_gaode.go              # 高德地图配置
│   ├── config_jwt.go                # JWT配置
│   ├── config_mysql.go              # MySQL配置
│   ├── config_qiniu.go              # 七牛云配置
│   ├── config_qq.go                 # QQ登录配置
│   ├── config_redis.go              # Redis配置
│   ├── config_system.go             # 系统配置
│   ├── config_upload.go             # 上传配置
│   ├── config_website.go            # 网站配置
│   ├── config_zap.go                # Zap日志配置
│   └── enter.go                     # 配置入口
├── core/                            # 核心功能
│   ├── config.go                    # 配置加载
│   ├── server.go                    # 服务器核心
│   ├── server_others.go             # 服务器其他功能
│   ├── server_windows.go            # Windows服务器功能
│   └── zap.go                       # 日志处理
├── flags/                           # 命令行工具与批量管理
│   ├── enter.go                     # 命令行入口
│   ├── flag_admin.go                # 管理相关命令
│   ├── flag_es.go                   # ES 相关命令
│   ├── flag_es_export.go            # ES 导出
│   ├── flag_es_import.go            # ES 导入
│   ├── flag_sql.go                  # SQL 相关命令
│   ├── flag_sql_export.go           # SQL 导出
│   └── flag_sql_import.go           # SQL 导入
├── global/                          # 全局对象
│   ├── global.go                    # 全局实例定义
│   └── model.go                     # 全局模型
├── initialize/                      # 初始化模块
│   ├── cron.go                      # 定时任务初始化
│   ├── es.go                        # ES 初始化
│   ├── gorm.go                      # 数据库初始化
│   ├── others.go                    # 其他初始化
│   ├── redis.go                     # Redis 初始化
│   └── router.go                    # 路由初始化
├── model/                           # 数据模型
│   ├── appTypes/                    # 应用类型
│   │   ├── image_category.go        # 图片分类
│   │   ├── image_storage.go         # 存储类型
│   │   ├── user_register.go         # 用户注册
│   │   └── user_role.go             # 用户角色
│   ├── database/                    # 数据库模型
│   │   ├── advertisement.go         # 广告
│   │   ├── artical_category.go      # 文章分类
│   │   ├── article_like.go          # 文章点赞
│   │   ├── article_tag.go           # 文章标签
│   │   ├── comment.go               # 评论
│   │   ├── feedback.go              # 反馈
│   │   ├── footer_link.go           # 底部链接
│   │   ├── friend_link.go           # 友情链接
│   │   ├── image.go                 # 图片
│   │   ├── jwt_blacklist.go         # JWT黑名单
│   │   ├── login.go                 # 登录
│   │   └── user.go                  # 用户
│   ├── elasticsearch/               # ES模型
│   │   └── article.go               # 文章
│   └── other/                       # 其他模型
│       └── es_index.go              # ES索引
├── request/                         # 请求参数模型
│   └── base.go                      # 基础请求模型
├── response/                        # 响应数据模型
│   ├── base.go                      # 域模型
│   └── response.go                  # 响应处理
├── router/                          # 路由配置
│   ├── base.go                      # 基础路由
│   └── enter.go                     # 路由入口
├── service/                         # 服务层
│   ├── base.go                      # 基础服务
│   ├── enter.go                     # 服务入口
│   └── es_index.go                  # ES索引服务
├── task/                            # 定时任务
│   └── enter.go                     # 任务入口
├── utils/                           # 工具函数
│   ├── email.go                     # 邮件工具
│   ├── hash.go                      # 密码加密
│   ├── parse.go                     # 解析工具
│   ├── random.go                    # 随机数工具
│   └── yaml.go                      # YAML工具
├── log/                             # 日志目录
├── config.yaml                      # 配置文件
└── main.go                          # 程序入口
```

### 数据模型

1. 用户系统
   - 用户管理与认证
   - 角色权限控制
   - JWT 认证
   - 第三方登录集成

2. 文章系统
   - 文章管理
   - 分类管理
   - 标签管理
   - 点赞系统
   - 评论系统

3. 其他功能
   - 友情链接
   - 底部链接
   - 广告管理
   - 反馈系统
   - 图片管理

### 核心功能

- 用户认证 (JWT)
- 验证码支持
- 邮件服务
- 文件上传（本地/七牛云）
- QQ第三方登录
- Redis缓存
- Elasticsearch搜索
- 高德地图服务集成
- Zap日志系统
- 命令行工具（批量导入导出、数据管理、运维脚本等）

### 技术栈

- 后端：Go
- ORM：Gorm
- 数据库：MySQL
- 缓存：Redis
- 搜索引擎：Elasticsearch
- 日志：Zap
- 任务调度：Cron
- 对象存储：七牛云

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