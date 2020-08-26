package houses

import (
	"github.com/gostarer/domain/services"
	"time"
)

type House struct {
	Id          int64     `db:id,omitempty`           //唯一编号
	HouseHid    string    `db:"house_hid",uni`        //楼盘唯一编号
	HouseName   string    `db:"house_name"`           //楼盘名称
	HouseType   int       `db:"house_type"`           //楼盘类型
	HouseStatus int       `db:"house_status"`         //楼盘状态
	CreatedAt   time.Time `db:"created_at,omitempty"` //创建时间
	UpdatedAt   time.Time `db:"updated_at,omitempty"` //更新时间
}

func (house *House) FromDTO(dto *services.HouseCreatedDTO) {
	house.HouseHid = dto.HouseId
	house.HouseName = dto.HouseName
	house.HouseType = dto.HouseType
	house.HouseStatus = dto.HouseStatus
}

func (house *House) ToDTO() *services.HouseDTO {
	dto := &services.HouseDTO{}
	dto.HouseId = house.HouseHid
	dto.HouseName = house.HouseName
	dto.HouseType = house.HouseType
	dto.HouseStatus = house.HouseStatus
	return dto
}
