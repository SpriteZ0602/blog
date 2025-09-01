# 个人博客

### 项目介绍

个人博客系统，前端使用Vue，后端使用go来构建

### 项目结构

``` 
├── server							# 后端
│   ├── config/						# 配置文件
│   │   └── config_{server}.go		# {server}的配置
│   ├── model/          			# 数据模型
│   │   └── appTypes         		# 链路追踪中间件
│   │   	└── image_storage.go	# 对图像存储位置（本地/云端）的判断
│	└── go.mod & go.sum          	# 后端Go模块依赖
└── README.md                		# 项目说明文档
```

