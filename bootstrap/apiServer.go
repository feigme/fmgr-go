package bootstrap

func StartApiServer() {
	// 初始化日志
	InitializeLog()

	// 初始化配置
	InitializeConfig()

	// 初始化db
	InitializeDB()

	// 启动服务器
	RunServer()
}
