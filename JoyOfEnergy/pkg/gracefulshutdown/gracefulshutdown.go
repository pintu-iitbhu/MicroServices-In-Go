package gracefulshutdown

import (
	"JoyOfEnergy/pkg/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	shutdownWaitGroup   *sync.WaitGroup
	shutdownChannel     chan int
	initVariableOnce    sync.Once
	exitChannelInstance chan struct{}
)

func Init(exitChannel chan struct{}) {
	if shutdownWaitGroup == nil {
		initVariableOnce.Do(func() {
			shutdownWaitGroup = new(sync.WaitGroup)
			shutdownChannel = make(chan int)
			exitChannelInstance = exitChannel
		})
	}
}

func ShutDown() {
	log := logger.NewLogger()

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quitChannel
	log.Info("Quit signal received ....")

	contextTimeOutInsecond := time.Duration(10)
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeOutInsecond*time.Second)
	log.Info("Quit signal received, sending shutdown and waiting on HTTP calls...")
	defer cancel()

	var srv http.Server
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Error Occurred which triggered shutdown", "Error", err)
	}

	log.Info("HTTP Server, shutdown gracefully.")

	log.Info("Quit signal received, sending shutdown and waiting on goroutines...")
	close(shutdownChannel)

	shutdownWaitGroup.Wait()
	log.Info("All go routines shutdown gracefully.")

	log.Info("main goroutine shutdown triggering...")
	close(exitChannelInstance)
}

func GetVariable() (*sync.WaitGroup, chan int) {
	return shutdownWaitGroup, shutdownChannel
}
