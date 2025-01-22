package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/a-korkin/ecommerce/configs"
	"github.com/a-korkin/ecommerce/internal/core/adapters/db"
	"github.com/a-korkin/ecommerce/internal/web/handlers"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var DBConnection *db.PostgresConnection
var KafkaProducer *kafka.Producer
var Server http.Server
var Ctx context.Context

func connectToDB() {
	var err error
	DBConnection, err = db.NewDBConnection(
		configs.GetEnv("GOOSE_DRIVER"), configs.GetEnv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
}

func connectToKafka() {
	var err error
	KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.GetEnv("KAFKA_HOST"),
	})
	if err != nil {
		log.Fatalf("failed to create kafka producer: %s", err)
	}
	log.Printf("kafka producer started")
}

func createWebServer() {
	Server = http.Server{
		Addr: ":8080",
	}
	router := handlers.NewRouter(DBConnection.DB, KafkaProducer)
	http.Handle("/", router)

	go func() {
		log.Println("server running")
		err := Server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
}

func runWebApp() {
	var stop context.CancelFunc
	Ctx, stop = signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	connectToDB()
	connectToKafka()
	createWebServer()

	<-Ctx.Done()
	log.Println("server terminated")
}

func shutDownWebApp() {
	log.Println("db connection closed")
	if err := DBConnection.CloseDBConnection(); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}

	log.Printf("kafka producer closed")
	KafkaProducer.Close()

	if err := Server.Shutdown(Ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
	log.Println("shutting down")
}

func runBrokerConsumer() {
	log.Printf("hello i love you, want you tell me your name?\n")
}

func shutDownBrokerConsumer() {
	log.Printf("goodbye my love, goodbye!\n")
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

	switch os.Args[1] {
	case "-w":
		fallthrough
	case "--web":
		runWebApp()
		defer shutDownWebApp()
	case "-b":
		fallthrough
	case "--broker":
		runBrokerConsumer()
		defer shutDownBrokerConsumer()
	default:
		usage()
	}
}
