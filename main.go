package main

import (
	"auth/model"
	"auth/router"
	"flag"
	"fmt"
)

var host, port, addr string

func init() {
	flag.StringVar(&host, "host", "localhost", "hostname")
	flag.StringVar(&port, "port", "8080", "port host")
	flag.Parse()
	addr = fmt.Sprintf("%s:%s", host, port)
	model.Migrations()
}

func main() {
	router.SetupServer().Run(addr)
	//listen_server := &http.Server{
	//	Addr: addr,
	//	Handler: r,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//listen_server.ListenAndServe()

}
