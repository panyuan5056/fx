package main

import (
	"fmt"
	"fx/pkg/setting"
	"fx/pkg/task"
	"fx/routers"
	"net/http"
)

func main() {
	task.Listen()
	fmt.Println("服务开始")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: setting.MaxHeaderBytes,
	}
	s.ListenAndServe()

}
