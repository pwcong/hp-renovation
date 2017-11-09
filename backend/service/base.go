package service

import (
	"github.com/pwcong/hp-renovation/backend/config"
	"github.com/pwcong/hp-renovation/backend/dao"
)

type BaseService struct {
	Conf *config.Config
	DAO  *dao.BaseDAO
}
