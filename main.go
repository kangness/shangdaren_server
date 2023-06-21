package main

import (
	"fmt"
	"github.com/kangness/shangdaren_server/config"
	"github.com/kangness/shangdaren_server/service"
	"net/http"
)

func main() {
	if err := config.InitServerConfig(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.HandlerHttpRequest)
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("error ", err)
		panic(err)
	}
}
