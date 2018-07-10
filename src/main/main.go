package main

import (
	"github.com/gin-gonic/gin"
	"apiserver/router"
	"log"
	"net/http"
	"time"
	"errors"
)

func main() {

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}


func pingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for thw router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}