package repository

import (
	"github.com/PonyWilliam/go-area/domain/model"
	"github.com/jinzhu/gorm"
)

type IArea interface {
	InitTable() error
	AddArea(area *model.Area)error
	UpdateArea(area *model.Area)error
	DelAreaByID(id int64)error
	FindAllArea() ([]model.Area,error)
}
func NewProductRepository(db *gorm.DB) IArea{
	return &AreaRepository{mysql: db}
}
type AreaRepository struct{
	mysql *gorm.DB
}
func(a *AreaRepository) InitTable() error{
	if(a.mysql.HasTable(&model.Area{})){
		a.mysql.DropTable(&model.Area{})
	}
	return a.mysql.CreateTable(&model.Area{}).Error
}
func(a *AreaRepository) AddArea(area *model.Area)error{
	return a.mysql.Model(area).Create(&area).Error
}
func(a *AreaRepository) UpdateArea(area *model.Area)error{
	return a.mysql.Model(&area).Where("id = ?",area.ID).Updates(area).Error
}
func(a *AreaRepository) DelAreaByID(id int64)error{
	return a.mysql.Where("id = ?",id).Delete(&model.Area{}).Error
}
func(a *AreaRepository) FindAllArea()(areas []model.Area,err error){
	return areas,a.mysql.Model(&areas).Find(&areas).Error
}