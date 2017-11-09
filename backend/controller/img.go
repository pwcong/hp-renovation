package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pwcong/hp-renovation/backend/service"
)

type ImgController struct {
	Base *BaseController
}

const (
	URL_IMG_UPLOAD = "/imgs"
	URL_IMG_GET    = "/imgs/:id"
)

func (self *ImgController) Upload(c echo.Context) error {

	form, err := c.MultipartForm()

	if err != nil {
		return c.JSON(http.StatusOK, BaseJSONResponse{
			Code: http.StatusBadRequest,
			Msg:  "上传文件不能为空",
		})
	}

	service := &service.ImgService{Base: self.Base.Service}
	var ids []int64

	imgs := form.File["imgs"]

	for _, img := range imgs {

		id, err := service.SaveImg(img)

		if err != nil {

			return c.JSON(http.StatusInternalServerError, BaseJSONResponse{
				Code: http.StatusInternalServerError,
				Msg:  err.Error(),
			})

		}

		ids = append(ids, id)

	}

	return c.JSON(http.StatusOK, BaseJSONResponse{
		Code: http.StatusOK,
		Msg:  "上传成功",
		Result: struct {
			IDS []int64 `json:"idx"`
		}{
			IDS: ids,
		},
	})
}

func (self *ImgController) Get(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, BaseJSONResponse{
			Code: http.StatusBadRequest,
			Msg:  "请求参数有误",
		})

	}

	service := &service.ImgService{Base: self.Base.Service}
	bytes, err := service.GetImg(int64(id))

	if err != nil {

		return c.JSON(http.StatusOK, BaseJSONResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})

	}

	c.Response().Header().Set("Cache-Control", "max-age=604800")
	return c.Blob(http.StatusOK, "image/png", bytes)

}
