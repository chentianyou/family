package main

import (
	"flag"
	"fmt"
	"github.com/chentianyou/family/family/route"
	"github.com/chentianyou/family/family/server"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"sync"
)

func main() {
	flag.StringVar(&server.ImageData,"img_data", "./data/simages.tar.gz", "the path of the images")
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover! error: %v", err)
			debug.PrintStack()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	//Starting REST Service.
	go func() {
		defer wg.Done()
		// setup router
		var RESTRouter = route.ServerRouter()
		var port = 80
		log.Println("server start as port 80")
		httpServer := &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: RESTRouter,
		}
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
			log.Println("httpServer send quit")
		}
	}()
	wg.Wait()
}
