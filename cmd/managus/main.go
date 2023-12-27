package main

import (
	"context"
	"flag"
	"github.com/IlyaZayats/managus/internal/db"
	"github.com/IlyaZayats/managus/internal/handlers"
	"github.com/IlyaZayats/managus/internal/repository"
	"github.com/IlyaZayats/managus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		dbUrl    string
		listen   string
		logLevel string
	)
	//postgres://dynus:dynus@localhost:5555/dynus
	//postgres://managus:managus@postgres_restus:5432/managus
	//postgres://managus:managus@postgres:5555/managus

	//postgres://managus:managus@localhost:5555/managus

	flag.StringVar(&dbUrl, "db", "postgres://managus:managus@postgres_managus:5432/managus", "database connection url")
	flag.StringVar(&listen, "listen", ":8080", "server listen interface")
	flag.StringVar(&logLevel, "log-level", "debug", "log level: panic, fatal, error, warning, info, debug, trace")

	flag.Parse()

	ctx := context.Background()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panicf("unable to get log level: %v", err)
	}
	logrus.SetLevel(level)

	dbc, err := db.NewPostgresPool(dbUrl)
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	managerRepository, err := repository.NewPostgresManagerRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build manager repo: %v", err)
	}

	clientRepository, err := repository.NewPostgresClientRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build client repo: %v", err)
	}

	invoiceRepository, err := repository.NewPostgresInvoiceRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build invoice repo: %v", err)
	}

	userRepo, err := repository.NewPostgresUserRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	managerService, err := services.NewManagerService(managerRepository)
	if err != nil {
		logrus.Panicf("unable build manager service: %v", err)
	}

	clientService, err := services.NewClientService(clientRepository)
	if err != nil {
		logrus.Panicf("unable build client service: %v", err)
	}

	invoiceService, err := services.NewInvoiceService(invoiceRepository)
	if err != nil {
		logrus.Panicf("unable build invoice service: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	g := gin.New()

	_, err = handlers.NewManagerHandlers(g, managerService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	_, err = handlers.NewClientHandlers(g, clientService)
	if err != nil {
		logrus.Panicf("unable build client handlers: %v", err)
	}

	_, err = handlers.NewInvoiceHandlers(g, invoiceService)
	if err != nil {
		logrus.Panicf("unable build invoice handlers: %v", err)
	}

	_, err = handlers.NewUserHandlers(g, userService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			}
		}
	}(childCtx)

	<-doneC

}
