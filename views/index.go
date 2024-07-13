package views

import (
	"errors"
	"hao/common"
	"hao/service"
	"log"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("首页获取数据出错: ", err)
		index.WriteError(w, errors.New("系统错误联系管理员"))
	}
	index.WriteDate(w, hr)
}
