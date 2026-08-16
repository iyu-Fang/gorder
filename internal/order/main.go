package main

import (
	"log"

	"github.com/iyu-Fang/gorder/common/config"
	"github.com/spf13/viper"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//log.Printf(viper.GetString("order.service-name"))
	log.Printf("%v", viper.Get("order"))
}
