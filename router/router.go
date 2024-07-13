package router

import (
	"hao/api"
	"hao/views"
	"net/http"
)

func index() {

}

func Router() {
	//1.页面 viewer。2.数据 api。3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
