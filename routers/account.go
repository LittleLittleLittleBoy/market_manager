package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/market_manager/pkg/ocr"
)

func identityAccount(c *gin.Context) {
	ocr.GetImageContent("http://20.245.143.120:8080/images/0f798c4d-c617-4de3-90bb-5ee052aa8f63.png")
}
