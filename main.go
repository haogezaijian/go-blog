package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "豪哥再见"
	indexData.Desc = "第一个go项目"
	t := template.New("index.html")
	//1.拿到当前的路劲
	path, _ := os.Getwd()
	//访问博客首页模版的时候，因为有多个模块的嵌套，解析文件的时候，需要将其涉及到的所有模版都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	//页面上涉及到的所有数据，
	t.Execute(w, indexData)
}

func main() {
	//程序入口，一个项目只能有一个入口

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
