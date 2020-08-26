package houses

import (
	"errors"
	"github.com/gostarer/domain/services"

	"github.com/segmentio/ksuid"
)

type houseDomain struct {
	house   House
	picture Picture
}

func NewHouseDomain() *houseDomain {
	return new(houseDomain)
}

func (domain *houseDomain) createPictureNo() {
	domain.picture.PictureNo = ksuid.New().Next().String()
}

func (domain *houseDomain) createHouseNo() {
	domain.house.HouseHid = ksuid.New().Next().String()
}

func (domain *houseDomain) CreateHouse(
	dto services.HouseCreatedDTO) (*services.HouseDTO, error) {

	var rdto *services.HouseDTO
	var id int64
	var err error
	domain.house = House{}
	domain.house.FromDTO(&dto)
	domain.createHouseNo()
	housedao := HouseDao{}

	if id, err = housedao.Insert(&domain.house); err != nil {
		return nil, err
	}
	if id < 0 {
		return nil, errors.New("创建楼盘数据失败！")
	}
	rdto = domain.house.ToDTO()
	return rdto, nil
}

func (domain *houseDomain) CreatePicture(
	dto services.PictureCreatedDTO) (*services.PictureDTO, error) {
	var id int64
	var err error
	var rdto *services.PictureDTO
	domain.picture = Picture{}
	domain.picture.FromDTO(&dto)
	domain.createPictureNo()
	picturedao := PictureDao{}
	if id, err = picturedao.Insert(&domain.picture); err != nil {
		return nil, err
	}
	if id < 0 {
		return nil, errors.New("创建图片失败！")
	}
	rdto = domain.picture.ToDTO()
	return rdto, nil
}
