package main

import (
	"example.com/andda/go-gin-example/pkg/setting"
	"example.com/andda/go-gin-example/routers"
	"fmt"
	"net/http"
)


func main(){


	router := routers.InitRouter()

	s:=&http.Server{
		Addr:           fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

   s.ListenAndServe()


}
