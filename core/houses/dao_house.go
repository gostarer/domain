package houses

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type HouseDao struct {
	runner *dbx.TxRunner
}

func (dao *HouseDao) GetOne(househid string) *House {
	house := &House{HouseHid: househid}
	ok, err := dao.runner.GetOne(house)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return house
}

func (dao *HouseDao) GetByUserId(
	id int,
	houseType int) *House {
	house := &House{}
	sql := "select * from house where id=? and house_type=? "
	ok, err := dao.runner.Get(house, sql, id, houseType)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return house
}

func (dao *HouseDao) Insert(house *House) (id int64, err error) {
	rs, err := dao.runner.Insert(house)
	if err != nil {
		return 0, nil
	}
	return rs.LastInsertId()
}

func (dao *HouseDao) UpdateStatus(
	houseHid string,
	status int,
) (rows int64, err error) {
	sql := "update house set status=? where house_hid=? "
	rs, err := dao.runner.Exec(sql, status, houseHid)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
