package api

import (
	"net/http"
	"strings"

	"bazi-analyzer/internal/service"

	"github.com/gin-gonic/gin"
)

type BaZiRequest struct {
	Bazi string `json:"bazi" form:"bazi" example:"壬戌 壬寅 庚午 丙戌" binding:"required"`
}

type BaZiResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// @Summary 八字分析接口（GET方式）
// @Description 输入八字四柱，返回完整6步分析结果。支持JSON和Markdown格式，通过Accept头或format参数控制
// @Tags 八字分析
// @Produce json,text/markdown
// @Param bazi query string true "八字四柱，格式：年柱 月柱 日柱 时柱" example(壬戌 壬寅 庚午 丙戌)
// @Param format query string false "返回格式：json或markdown" Enums(json,markdown) default(json)
// @Success 200 {object} api.BaZiResponse
// @Failure 400 {object} api.BaZiResponse
// @Router /bazi/analyze [get]
func AnalyzeHandlerGet(c *gin.Context) {
	AnalyzeHandler(c)
}

// @Summary 八字分析接口（POST方式）
// @Description 输入八字四柱，返回完整6步分析结果。支持JSON和Markdown格式，通过Accept头或format参数控制
// @Tags 八字分析
// @Accept json
// @Produce json,text/markdown
// @Param request body api.BaZiRequest true "八字请求体"
// @Param format query string false "返回格式：json或markdown" Enums(json,markdown) default(json)
// @Success 200 {object} api.BaZiResponse
// @Failure 400 {object} api.BaZiResponse
// @Router /bazi/analyze [post]
func AnalyzeHandlerPost(c *gin.Context) {
	AnalyzeHandler(c)
}

func AnalyzeHandler(c *gin.Context) {
	var req BaZiRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaZiResponse{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	result, err := service.AnalyzeBaZi(req.Bazi)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaZiResponse{
			Code:    400,
			Message: "八字解析失败: " + err.Error(),
		})
		return
	}

	accept := c.GetHeader("Accept")
	format := c.Query("format")

	if strings.Contains(accept, "text/markdown") || format == "markdown" {
		c.Header("Content-Type", "text/markdown; charset=utf-8")
		c.String(http.StatusOK, result.ToMarkdown())
		return
	}

	c.JSON(http.StatusOK, BaZiResponse{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}
