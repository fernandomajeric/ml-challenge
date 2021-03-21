package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/fernandomajeric/ml-challenge/app"
	"github.com/fernandomajeric/ml-challenge/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	var configFilePath string
	var serverPort = getEnv("PORT","8080")
	var redisInstance = getEnv("REDIS_URL","localhost:6379")

	flag.StringVar(&configFilePath, "config", "./", "absolute path to the configuration file")
	flag.StringVar(&serverPort, "server_port", serverPort, "port on which server runs")
	flag.Parse()

	application := app.New(configFilePath)

	//Test Db
	config.Configuration.RedisCache.Addrs = getEnv("REDIS_URL", config.Configuration.RedisCache.Addrs)

	client := redis.NewClient(&redis.Options{
		Addr:     redisInstance,
		Password: config.Configuration.RedisCache.Password, // no password set
		DB:       config.Configuration.RedisCache.DB,       // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	if err != nil {
		panic(err)
	}

	// start http server
	application.Start(serverPort)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
