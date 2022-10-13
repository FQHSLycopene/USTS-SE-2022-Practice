package main

import (
	_ "BackEnd/docs"
	"BackEnd/router"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//func init() {
//}

func main() {
	r := router.Router()
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
