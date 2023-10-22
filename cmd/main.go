package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"test_shop/pkg/handler"
	repository "test_shop/pkg/repositry"
	"test_shop/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("you have error %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("you have error with dotenv file %s", err.Error())
	}
	db, _ := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	logrus.Print("Shop Started")
	_, err := handlers.InitConsole(4)
	if err != nil {
		fmt.Println(err)
		logrus.Fatalf("you have error %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
