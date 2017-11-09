package service

import (
	"errors"
	"io/ioutil"
	"mime/multipart"

	"github.com/pwcong/hp-renovation/backend/dao"
)

type ImgService struct {
	Base *BaseService
}

func (self *ImgService) SaveImg(img *multipart.FileHeader) (int64, error) {

	src, err := img.Open()
	if err != nil {
		return 0, errors.New("文件读取失败")
	}
	defer src.Close()

	bytes, err := ioutil.ReadAll(src)
	if err != nil {
		return 0, errors.New("文件读取失败")
	}

	ImgDAO := dao.ImgDAO{Base: self.Base.DAO}

	return ImgDAO.Add(bytes)
}

func (self *ImgService) GetImg(id int64) ([]byte, error) {

	ImgDAO := dao.ImgDAO{Base: self.Base.DAO}
	res, err := ImgDAO.Get(id)
	if err != nil {
		return []byte{}, err
	}

	return res.Src, err

}
