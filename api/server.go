package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	/*
		为什么这块使用指针而不是直接传入对象：
		1. 避免在函数见传递结构体时进行复制，也可以直接对结构体进行修改
		2. go一般都使用指针来创建结构体实例
	*/
	server := &Server{store: store}
	/*
		gin.default做了什么：
		1. 创建一个gin的引擎实例
		2. 自动添加中间件：logger（记录http请求日志），recovery（捕捉panic，防止服务器崩溃）
	*/
	router := gin.Default()

	// 注册路由： 将POST /accounts请求交给server.createAccount方法处理
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
