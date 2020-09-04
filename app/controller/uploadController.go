package controller

import (
	"github.com/gin-gonic/gin"
	"golangPro/golang-mallapi/pkg/logging"
	"golangPro/golang-mallapi/pkg/status"
	"golangPro/golang-mallapi/pkg/upload"
	"golangPro/golang-mallapi/pkg/util"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]

func UploadImage(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}
	file, image, err := context.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		utilGin.Response(status.ERROR, status.GetMsg(status.ERROR), nil)
		return
	}

	if image == nil {
		utilGin.Response(status.INVALID_PARAMS, status.GetMsg(status.INVALID_PARAMS), nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		utilGin.Response(status.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, status.GetMsg(status.ERROR_UPLOAD_CHECK_IMAGE_FORMAT), nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		utilGin.Response(status.ERROR_UPLOAD_CHECK_IMAGE_FAIL, status.GetMsg(status.ERROR_UPLOAD_CHECK_IMAGE_FAIL), nil)
		return
	}

	if err := context.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		utilGin.Response(status.ERROR_UPLOAD_SAVE_IMAGE_FAIL, status.GetMsg(status.ERROR_UPLOAD_SAVE_IMAGE_FAIL), nil)
		return
	}

	utilGin.Response(status.SUCCESS, status.GetMsg(status.SUCCESS), map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})

}
