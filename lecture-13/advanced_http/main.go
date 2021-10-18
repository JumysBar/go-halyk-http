package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
	"golang.org/x/net/http2"
)

func main() {
	var httpServer = http.Server{
		Addr: ":9191",
	}
	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&httpServer, &http2Server)
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost:9191/hello/sayHello")
	go log.Fatal(httpServer.ListenAndServeTLS("./cert/server.crt", "./cert/server.key"))

	//HTTP 3.0
	server := http3.Server{
		Server: &http.Server{
			Addr:    "localhost:8081",
			Handler: http.HandlerFunc(echoPayload),
		},
	}

	log.Fatalln(server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key"))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
