package dao

import (
	"time"

	"github.com/pwcong/hp-renovation/backend/model"
)

type ImgDAO struct {
	Base *BaseDAO
}

const SCHEMA_ADD = `INSERT INTO imgs (src, createdAt) VALUES (?, ?)`
const SCHEMA_GET = `SELECT src FROM imgs WHERE id=?`

func (self *ImgDAO) Add(src []byte) (int64, error) {

	res, err := self.Base.DB.MySQL.Exec(SCHEMA_ADD, src, time.Now())

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (self *ImgDAO) Get(id int64) (model.Img, error) {

	var res model.Img

	err := self.Base.DB.MySQL.QueryRowx(SCHEMA_GET, id).StructScan(&res)

	if err != nil {
		return model.Img{}, err
	}

	return res, nil
}
