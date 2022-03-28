package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/rodrigojb/meli-proxy/meli-proxy/cmd/handler"
	"github.com/rodrigojb/meli-proxy/meli-proxy/config"
	"github.com/rodrigojb/meli-proxy/meli-proxy/internal"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	rdb, err := newRedisDB()
	if err != nil {
		return fmt.Errorf("initiating redis db: %v", err)
	}

	rpURL, err := url.Parse("https://api.mercadolibre.com")
	if err != nil {
		return fmt.Errorf("parsing url: %v", err)
	}

	handler := handler.Handler(rpURL, rdb, criteria())

	fmt.Println("Starting MELI-PROXY on port " + config.ServerPort)

	return http.ListenAndServe(":"+config.ServerPort, handler)
}

func newRedisDB() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("connecting to redis: %v", err)
	}

	return rdb, nil
}

func criteria() []internal.Criterion {
	return []internal.Criterion{
		internal.IpCriterion{
			IP:     "localhost",
			Limit:  1500,
			Period: 60,
		},
		internal.PathCriterion{
			Path:   "/MLA3530",
			Limit:  1000,
			Period: 60,
		},
		internal.PathCriterion{
			Path:   "/MLA3531",
			Limit:  1000,
			Period: 60,
		},
	}
}
