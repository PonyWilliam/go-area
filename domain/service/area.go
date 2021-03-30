package service

import (
	"github.com/PonyWilliam/go-area/domain/model"
	"github.com/PonyWilliam/go-area/domain/repository"
)

type IAreaServices interface {
	AddArea(area *model.Area)error
	DelAreaByID(id int64)error
	FindAllArea() ([]model.Area,error)
}
func NewAreaServices(area repository.IArea)IAreaServices{
	return &AreaServices{areaRepository: area}
}
type AreaServices struct{
	areaRepository repository.IArea
}
func(a *AreaServices)AddArea(area *model.Area)error{
	return a.areaRepository.AddArea(area)
}
func(a *AreaServices)DelAreaByID(id int64)error{
	return a.areaRepository.DelAreaByID(id)
}
func(a *AreaServices)FindAllArea()([]model.Area,error){
	return a.areaRepository.FindAllArea()
}

