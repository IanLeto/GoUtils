package utils

import (
	"fmt"
	"net/http"
)

type HttpServiceOption struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Port    string `json:"port"`
	handler func(w http.ResponseWriter, req *http.Request)
}

func NewHttpServer(opt *HttpServiceOption) {
	if opt.handler != nil {
		http.HandleFunc("/hello", opt.handler)
	} else {
		http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println(1111)
			fmt.Fprintf(writer, "world\n")
		})
	}
	NoErr(http.ListenAndServe(opt.Addr+":"+opt.Port, nil))
}

func RunHttpService() {
	http.HandleFunc("/hello2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`world`))
	})
	http.ListenAndServe("http://localhost:4999", nil)

}
