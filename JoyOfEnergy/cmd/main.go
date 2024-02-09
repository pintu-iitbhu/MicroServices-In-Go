package main

import (
	"JoyOfEnergy/internal/api/server"
	"JoyOfEnergy/internal/errorhandling"
	"JoyOfEnergy/pkg/gracefulshutdown"
	"JoyOfEnergy/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	log.Info("welcome to joy of energy service", "Joy", "energey")

	exitChannel := make(chan struct{})
	//flag.Usage = func() {
	//	fmt.Println("Usage: server -s {service_name} -e {environment}")
	//	os.Exit(1)
	//}
	//flag.Parse()

	gracefulshutdown.Init(exitChannel)
	go server.Init()
	errorhandling.NewErrorStatusMessage()
	errorhandling.NewHttpStatusMap()

	gracefulshutdown.ShutDown()
	<-exitChannel
	log.Info("main goroutine shutdown completed gracefully.")
}
