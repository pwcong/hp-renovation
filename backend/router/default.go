package router

import (
	"github.com/pwcong/hp-renovation/backend/config"
	"github.com/pwcong/hp-renovation/backend/controller"
	"github.com/pwcong/hp-renovation/backend/dao"
	"github.com/pwcong/hp-renovation/backend/db"
	"github.com/pwcong/hp-renovation/backend/service"

	"github.com/labstack/echo"
)

const API_VERSION = "v1"

const GLOBAL_PREFIX = "/api/" + API_VERSION

func Init(e *echo.Echo, conf *config.Config, db *db.DB) {

	e.Static("/", "view")

	baseDAO := &dao.BaseDAO{Conf: conf, DB: db}
	baseService := &service.BaseService{Conf: conf, DAO: baseDAO}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	imgController := &controller.ImgController{Base: baseController}

	e.GET(GLOBAL_PREFIX+controller.URL_IMG_GET, imgController.Get)

	e.POST(GLOBAL_PREFIX+controller.URL_IMG_UPLOAD, imgController.Upload)

}
