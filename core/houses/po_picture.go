package houses

import (
	"github.com/gostarer/domain/services"
	"time"
)

type Picture struct {
	Id            int64     `db:id,omitempty`           //唯一编号
	PictureNo     string    `db:"picture_id",uni`       //图片唯一编号
	PictureUrl    string    `db:"picture_url"`          //图片路径
	PictureType   int       `db:"picture_type"`         //图片类型
	PictureStatus int       `db:"picture_status"`       //图片状态
	CreatedAt     time.Time `db:"created_at,omitempty"` //创建时间
	UpdatedAt     time.Time `db:"updated_at,omitempty"` //更新时间
}

func (picture *Picture) FromDTO(dto *services.PictureCreatedDTO) {
	picture.PictureNo = dto.PictureNo
	picture.PictureUrl = dto.PictureUrl
	picture.PictureType = dto.PictureType
	picture.PictureStatus = dto.PictureStatus
}

func (picture *Picture) ToDTO() *services.PictureDTO {
	dto := &services.PictureDTO{}
	dto.PictureId = picture.PictureNo
	dto.PictureUrl = picture.PictureUrl
	dto.PictureType = picture.PictureType
	dto.PictureStatus = picture.PictureStatus
	return dto
}
