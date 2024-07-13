package common

import (
	"hao/config"
	"hao/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(any(err))
		}
		w.Done()
	}()
	w.Wait()
}
