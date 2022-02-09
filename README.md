# 项目介绍

## 规范
### 命名规范
> 参考https://www.cnblogs.com/rickiyang/p/11074174.html

1. 文件命名
这点有点不同，建议采用习惯的驼峰规则命名

2. 单元测试
GoStub+GoConvey+GoMock

go get github.com/smartystreets/goconvey

## 目录结构
app/common公共模块（请求、响应结构体等）
app/controllers业务调度器
app/middleware中间件
app/models数据库结构体
app/services业务层
bootstrap项目启动初始化
config配置结构体
global全局变量
routes路由定义
static静态资源（允许外部访问）
storage系统日志、文件等静态资源）
utils工具函数
config.yaml配置文件
main.go项目启动文件

## 包选择
- gorm
- zap
- excel：github.com/xuri/excelize/v2, 文档https://xuri.me/excelize/zh-hans/