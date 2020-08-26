package services

import "github.com/gostarer/domain/infra/base"

var IHouseService HouseService

func GetHouseService() HouseService {
	base.Check(IHouseService)
	return IHouseService
}

type HouseService interface {
	CreateHouse(dto HouseCreatedDTO) (*HouseDTO, error)
	CreatePicture(dto PictureCreatedDTO) (*PictureDTO, error)
	GetHouseByHid(houseid string) *HouseDTO
	GetPictureByNo(pictureNo string) *PictureDTO
}

type HouseCreatedDTO struct {
	HouseId     string `validate:"required"`
	HouseName   string `validate:"required"`
	HouseType   int
	HouseStatus int
}

type PictureCreatedDTO struct {
	PictureNo     string `validate:"required"`
	PictureUrl    string `validate:"required"`
	PictureType   int
	PictureStatus int
}

type HouseDTO struct {
	HouseId     string `validate:"required"`
	HouseName   string `validate:"required"`
	HouseType   int
	HouseStatus int
}

type PictureDTO struct {
	PictureId     string `validate:"required"`
	PictureUrl    string `validate:"required"`
	PictureType   int
	PictureStatus int
}
