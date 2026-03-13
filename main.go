package main

import (
	"net/http"

	"bazi-analyzer/internal/api"

	_ "bazi-analyzer/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 八字分析系统 API
// @version 1.0
// @description 基于Go语言开发的八字分析系统，提供完整的6步分析流程
// @BasePath /api/v1

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "八字分析系统 API Server\n访问 /swagger/index.html 查看接口文档")
	})

	r.GET("/api/v1/bazi/analyze", api.AnalyzeHandlerGet)
	r.POST("/api/v1/bazi/analyze", api.AnalyzeHandlerPost)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
