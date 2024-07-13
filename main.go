package main

import (
	"hao/common"
	"hao/router"
	"log"
	"net/http"
)

func init() {
	//模版加载
	common.LoadTemplate()
}

func main() {
	//程序入口，一个项目只能有一个入口

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
