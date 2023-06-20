package main

import (
	"fmt"
	"github.com/kangness/shangdaren_server/db"
	"github.com/kangness/shangdaren_server/service"
	"log"
	"net/http"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
