package handler

import (
	"context"
	"github.com/PonyWilliam/go-area/domain/model"
	"github.com/PonyWilliam/go-area/domain/service"
	area "github.com/PonyWilliam/go-area/proto"
)
type Area struct{
	AreaService service.IAreaServices
}
func(a *Area)CreateArea(ctx context.Context,req *area.Request_Add_Area,response *area.ResponseMessage) error{
	Area := &model.Area{
		Name: req.Name,
		Description: req.Description,
	}
	err := a.AreaService.AddArea(Area)
	if err!=nil{
		return err
	}
	response.Message = "success"
	return nil
}
func(a *Area)UpdateArea(ctx context.Context,req *area.Request_Update_Area,response *area.ResponseMessage) error{
	Area := &model.Area{
		ID: req.Id,
		Name: req.Name,
		Description: req.Description,
	}
	err := a.AreaService.UpdateArea(Area)
	if err!=nil{
		return err
	}
	response.Message = "success"
	return nil
}
func(a *Area)DelArea(ctx context.Context,req *area.Request_AreaID,response *area.ResponseMessage) error{
	err := a.AreaService.DelAreaByID(req.Id)
	if err != nil{
		return err
	}
	response.Message = "success"
	return nil
}
func(a *Area)FindAll(ctx context.Context,req *area.Request_NULL,response *area.Response_AreaInfos) error{
	rsp,err := a.AreaService.FindAllArea()
	if err != nil{
		response.Infos = nil
		return err
	}
	for _,v := range rsp{
		temp := &area.Response_AreaInfo{}
		Swap(v,temp)
		response.Infos = append(response.Infos,temp)
	}
	return nil
}

func Swap(req model.Area,rsp *area.Response_AreaInfo){
	rsp.Name = req.Name
	rsp.Id = req.ID
	rsp.Description = req.Description
}