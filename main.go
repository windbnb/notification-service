package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"github.com/windbnb/notification-service/handler"
	repository "github.com/windbnb/notification-service/repository"
	"github.com/windbnb/notification-service/router"
	"github.com/windbnb/notification-service/service"
	"github.com/windbnb/notification-service/tracer"
	"github.com/windbnb/notification-service/util"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	db := util.ConnectToMongoDatabase()

	tracer, closer := tracer.Init("notification-service")
	opentracing.SetGlobalTracer(tracer)
	router := router.ConfigureRouter(&handler.Handler{
		Tracer: tracer,
		Closer: closer,
		Service: &service.RatingService{
			Repo: &repository.Repository{
				Db: db}}})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3005"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	servicePath, servicePathFound := os.LookupEnv("SERVICE_PATH")
	if !servicePathFound {
		servicePath = "localhost:8085"
	}

	srv := &http.Server{Addr: servicePath, Handler: c.Handler(router)}

	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
