package main

import (
	"context"
	"github.com/BlackRRR/shortened-links/internal/app/handler"
	"github.com/BlackRRR/shortened-links/internal/app/repository"
	"github.com/BlackRRR/shortened-links/internal/app/services"
	"github.com/BlackRRR/shortened-links/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.TODO()

	//Initializing config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
	}

	log.Println("init config success")

	dataBase, err := pgxpool.ConnectConfig(ctx, cfg.RepositoryCfg)
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

	handlers := handler.NewHandler(svc)

	log.Println("init handlers success")

	go func() {
		log.Println("http server started on port:", cfg.ServicePort)
		serviceErr := http.ListenAndServe(":"+cfg.ServicePort, handlers.InitRoutes())
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
