# 项目介绍

## 规范
### 命名规范
> 参考https://www.cnblogs.com/rickiyang/p/11074174.html
1. 包名称
保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为小写单词，不要使用下划线或者混合大小写
2. 文件命名
尽量采取有意义的文件名，简短，有意义，应该为小写单词，~~使用下划线分隔各个单词~~，
建议采用习惯的驼峰规则命名

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
