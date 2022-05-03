package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"log"
	"module_30/pkg/storage"
	"os"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initialization: %s", err.Error())
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", viper.Get("username"), viper.Get("password"), viper.Get("host"), viper.Get("port"), viper.Get("database"))
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	db, err := storage.New(dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database NEWWW: %v\n", err)
		os.Exit(1)
	}

	example, err := db.AuthorsTasks(1)
	fmt.Println(example, err)
	example2, err := db.Task(10)
	fmt.Println(example2, err)

}
