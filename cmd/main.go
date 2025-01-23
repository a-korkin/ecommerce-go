package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/a-korkin/ecommerce/configs"
	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/core/services"
	"github.com/a-korkin/ecommerce/internal/web/handlers"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// var DBConnection *db.PostgresConnection
// var KafkaProducer *kafka.Producer
// var Server http.Server
// var Ctx context.Context

type AppState struct {
	DBConnection  *db.PostgresConnection
	KafkaProducer *kafka.Producer
	Server        http.Server
	KafkaService  *services.OrderService
	Ctx           context.Context
}

func NewAppState() *AppState {
	return &AppState{}
}

var appState *AppState

func connectToDB() {
	var err error
	appState.DBConnection, err = db.NewDBConnection(
		configs.GetEnv("GOOSE_DRIVER"), configs.GetEnv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
}

func connectToKafka() {
	var err error
	appState.KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.GetEnv("KAFKA_HOST"),
	})
	if err != nil {
		log.Fatalf("failed to create kafka producer: %s", err)
	}
	log.Printf("producer started")
}

func createWebServer() {
	appState.Server = http.Server{
		Addr: ":8080",
	}
	router := handlers.NewRouter(appState.DBConnection.DB, appState.KafkaProducer)
	http.Handle("/", router)

	go func() {
		log.Println("server running")
		err := appState.Server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
}

func runWebApp() {
	var stop context.CancelFunc
	appState.Ctx, stop = signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	connectToDB()
	connectToKafka()
	createWebServer()
	defer shutDownWebApp()

	<-appState.Ctx.Done()
	log.Println("server terminated")
}

func shutDownWebApp() {
	log.Println("db connection closed")
	if err := appState.DBConnection.CloseDBConnection(); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}

	log.Printf("producer closed")
	appState.KafkaProducer.Close()

	if err := appState.Server.Shutdown(appState.Ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
	log.Println("shutting down")
}

func runBrokerConsumer() {
	var stop context.CancelFunc
	appState.Ctx, stop = signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	connectToDB()

	appState.KafkaService = services.NewOrderService(
		appState.DBConnection.DB,
		configs.GetEnv("KAFKA_HOST"),
		configs.GetEnv("KAFKA_TOPIC"))

	defer shutDownBrokerConsumer()
	log.Printf("consumer started")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go appState.KafkaService.Run(appState.Ctx, &wg)
	wg.Wait()
}

func shutDownBrokerConsumer() {
	if !appState.KafkaService.Consumer.IsClosed() {
		appState.KafkaService.ShutDown()
	}

	log.Println("db connection closed")
	if err := appState.DBConnection.CloseDBConnection(); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}
	log.Printf("consumer stoped")
}

func usage() {
	fmt.Printf("Usage: make run [OPTION]\n")
	fmt.Printf("	-w, --web		Run web server\n")
	fmt.Printf("	-b, --broker	Run message broker(kafka)\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments\n\n")
		usage()
		os.Exit(1)
	}

	appState = NewAppState()

	switch os.Args[1] {
	case "-w":
		fallthrough
	case "--web":
		runWebApp()
	case "-b":
		fallthrough
	case "--broker":
		runBrokerConsumer()
	default:
		usage()
	}
}
