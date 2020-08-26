package houses

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type PictureDao struct {
	runner *dbx.TxRunner
}

func (dao *PictureDao) GetOne(picNo string) *Picture {
	picture := &Picture{PictureNo: picNo}
	ok, err := dao.runner.GetOne(picture)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return picture
}

func (dao *PictureDao) GetByPicType(picType int) *Picture {
	out := &Picture{}
	sql := "select * from picture where pic_type=? "
	ok, err := dao.runner.Get(out, sql, picType)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return out
}

func (dao *PictureDao) Insert(p *Picture) (id int64, err error) {
	rs, err := dao.runner.Insert(p)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}
