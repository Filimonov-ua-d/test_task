package main

import (
	"log"

	"github.com/Filimonov-ua-d/test_task/config"
	"github.com/Filimonov-ua-d/test_task/server"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
