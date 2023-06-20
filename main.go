package main

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

const Port string = ":8080"

// HandlerRequest 处理http请求
func HandlerRequest(resp http.ResponseWriter, req *http.Request) {
	if req != nil && req.Method == "HEAD" {
		resp.Write([]byte("OK"))
		return

	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Info("request url:", req.URL.Path)
	glog.Info("request body: ", string(body))
}

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", HandlerRequest)
	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Println(err)
	}
}
