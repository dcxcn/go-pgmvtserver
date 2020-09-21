package router

import (
	"go-pgmvtserver/app/controller"

	"github.com/kataras/iris/v12"
)

// Router iris的
func Router(irisApp *iris.Application) {
	namespace := "/pg_postgis_mvt/api/v1"
	// 为了看起来更容易理解
	router := irisApp
	router.Get(namespace+"/{srcLayerName:string}/{tableName:string}/{z:int}/{x:int}/{y:int}", controller.Mvt.GetMvt)
}
