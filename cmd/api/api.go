package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meetup/interals/store"
)

type App struct {
	config Config
	store  store.Storage
}

type Config struct {
	addr string
}

func (a *App) mount() http.Handler {
	/*
		   //new mux
			mux := http.NewServeMux()
			//end points
			mux.HandleFunc("GET /v1/health", a.health)
			return mux
	*/

	//changing to Gin
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/health", a.healthCheckHandler)

	}

	return r
}

func (a *App) run(mux http.Handler) error {

	//mux config
	srv := &http.Server{
		Addr:         a.config.addr,    // port no
		Handler:      mux,              //new mux route handler
		WriteTimeout: time.Second * 30, //max time to write to response
		ReadTimeout:  time.Second * 10, // max time to read from request
		IdleTimeout:  time.Minute,
	}
	//listning on port 8080
	defer srv.Close()

	log.Printf("Server started at %s", a.config.addr)

	return srv.ListenAndServe()
}

func (a *App) healthCheckHandler(ctx *gin.Context) {

	res := map[string]string{"message": "ok"}

	ctx.JSON(200, res)
}
