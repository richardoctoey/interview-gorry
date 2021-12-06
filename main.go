package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"richardoctoey/interview-gorry/api"
	"richardoctoey/interview-gorry/common"
	"richardoctoey/interview-gorry/event"
	"syscall"
)

func main() {
	godotenv.Load("configuration/config.env")
	err := common.StartDatabase(
		os.Getenv("dbuser"),
		os.Getenv("dbpass"),
		os.Getenv("dbname"),
		os.Getenv("dbhost"),
		os.Getenv("dbport"))
	if err != nil {
		panic(err)
	}
	common.AutoMigrate(&event.Event{})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Cleaning up..") //add cleanup here
		os.Exit(1)
	}()

	api.StartApi(fmt.Sprintf("%s:%s", os.Getenv("host"), os.Getenv("port")))
}
