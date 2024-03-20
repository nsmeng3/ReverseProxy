package main

import (
	"log"
	"main/cmd"
	"main/handle"
	"net/http"
)

var srv http.Server

func StartServer(bind string, remote string) {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	h := &handle.Handle{ReverseProxy: remote}
	srv.Addr = bind
	srv.Handler = h
	//go func() {
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
	//}()
}

func StopServer() {
	if err := srv.Shutdown(nil); err != nil {
		log.Println(err)
	}
}

func main() {
	parseCmd := cmd.ParseCmd()
	StartServer(parseCmd.Bind, parseCmd.Remote)
}
