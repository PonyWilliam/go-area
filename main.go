package main

import (
	"github.com/PonyWilliam/go-area/domain/repository"
	services2 "github.com/PonyWilliam/go-area/domain/service"
	"github.com/PonyWilliam/go-area/handler"
	area2 "github.com/PonyWilliam/go-area/proto"
	common "github.com/PonyWilliam/go-common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/opentracing/opentracing-go"
	"strconv"
	"time"
)
func main() {
	consulConfig,err := common.GetConsualConfig("1.116.62.214",8500,"/micro/config")
	//配置中心
	if err != nil{
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"1.116.62.214"}
			options.Timeout = time.Second * 10
		},
	)
	t,io,err := common.NewTracer("go.micro.service.area",":6833")
	if err != nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// Create service

	srv := micro.NewService(
		micro.Name("go.micro.service.area"),
		micro.Version("latest"),
		micro.Address(":8091"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(common.QPS)),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	go common.PrometheusBoot("5004")
	rp := repository.NewProductRepository(db)
	_ = rp.InitTable()

	areaServices := services2.NewAreaServices(repository.NewProductRepository(db))
	err = area2.RegisterAreaHandler(srv.Server(),&handler.Area{AreaService: areaServices})

	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}