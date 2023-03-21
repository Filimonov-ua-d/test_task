package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
}

func NewApp(db *sqlx.DB) *App {

	var err error

	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")
	sslmode := viper.GetString("postgres.sslmode")
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		user, password, dbname, sslmode, host, port)

	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		log.Panic(err)
	}

	return &App{}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("running server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
