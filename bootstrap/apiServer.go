package bootstrap

func StartApiServer() {

	// 初始化db
	InitializeDB()

	// 启动服务器
	RunServer()
}
