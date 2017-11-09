package dao

import (
	"github.com/pwcong/hp-renovation/backend/config"
	"github.com/pwcong/hp-renovation/backend/db"
)

type BaseDAO struct {
	Conf *config.Config
	DB   *db.DB
}
