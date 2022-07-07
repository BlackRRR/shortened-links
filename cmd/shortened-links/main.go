package main

import (
	"context"
	"github.com/BlackRRR/shortened-Links/internal/app/repository"
	"github.com/BlackRRR/shortened-Links/internal/app/server"
	"github.com/BlackRRR/shortened-Links/internal/app/services"
	"github.com/BlackRRR/shortened-Links/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.TODO()

	//Initializing config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}

	log.Println("init config success")

	dataBase, err := repository.InitDataBase(ctx, cfg.RepositoryCfg)
	if err != nil {
		log.Fatalf("failed to init database: %s", err.Error())
	}

	log.Println("init database success")

	repositories, err := repository.InitRepositories(ctx, dataBase)
	if err != nil {
		log.Fatalf("failed to init repositories: %s", err.Error())
	}

	log.Println("init repositories success")

	svc := services.InitServices(repositories)

	log.Println("init services success")

	httpHandler := server.MakeHTTPHandler(server.NewServer(svc))

	go func() {

		log.Println("http server started on port:", cfg.ServicePort)
		serviceErr := http.ListenAndServe(":"+cfg.ServicePort, httpHandler)
		if serviceErr != nil {
			log.Fatalf("http handler was stoped by err: %s", serviceErr.Error())
		}
	}()

	sig := <-subscribeToSystemSignals()

	log.Printf("shutdown all process on '%s' system signal\n", sig.String())

}

func subscribeToSystemSignals() chan os.Signal {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	return ch
}