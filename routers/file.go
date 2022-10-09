package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/littlelittlelittleboy/market_manager/pkg/app"
	"github.com/littlelittlelittleboy/market_manager/pkg/e"
	"github.com/littlelittlelittleboy/market_manager/pkg/upload"
)

func uploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	_, image, err := c.Request.FormFile("image")

	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	isAllowed := upload.CheckImageExt(image.Filename)
	if !isAllowed {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	filePath := upload.GetTempFilePath()
	if err := c.SaveUploadedFile(image, filePath); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url": filePath,
	})
}
