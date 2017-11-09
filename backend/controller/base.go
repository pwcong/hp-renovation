package controller

import (
	"github.com/pwcong/hp-renovation/backend/config"
	"github.com/pwcong/hp-renovation/backend/service"
)

type BaseController struct {
	Conf    *config.Config
	Service *service.BaseService
}

type BaseJSONResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}
