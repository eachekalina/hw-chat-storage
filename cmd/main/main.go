package main

import (
	"context"
	"fmt"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/grpc"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/handler"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/kafka"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/redis"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/repository"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	redislib "github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	err := run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	db struct {
		connString string
	}
	redis struct {
		addr string
	}
	kafka struct {
		brokers []string
		group   string
		topics  []string
	}
}

func run(ctx context.Context) error {
	cfg, err := getConfig()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.db.connString)
	if err != nil {
		return fmt.Errorf("pgxpool: %w", err)
	}
	defer pool.Close()

	msgRepo := repository.NewMessageRepository(pool)
	userRepo := repository.NewUserRepository(pool)

	redisCli := redislib.NewClient(&redislib.Options{
		Addr: cfg.redis.addr,
	})
	cache := redis.NewCache(redisCli)

	msgSvc := service.NewMessageService(msgRepo, cache)
	userSvc := service.NewUserService(userRepo)

	h := handler.NewHandler(msgSvc)

	grpcMsgHandler := grpc.NewMessageHandler(msgSvc)
	grpcUserHandler := grpc.NewUserHandler(userSvc)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		err = kafka.Run(ctx, h.HandleMessage, cfg.kafka.brokers, cfg.kafka.group, cfg.kafka.topics)
		if err != nil {
			return fmt.Errorf("kafka: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		err := grpc.Run(ctx, grpcMsgHandler, grpcUserHandler)
		if err != nil {
			return fmt.Errorf("grpc: %w", err)
		}
		return nil
	})

	return eg.Wait()
}

func getConfig() (config, error) {
	cfg := config{}
	cfg.db.connString = lookupEnvDefault("DB_URL", "postgres://postgres:mysecretpassword@localhost:5432/postgres")
	cfg.redis.addr = lookupEnvDefault("REDIS_ADDR", "localhost:6379")
	cfg.kafka.brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	cfg.kafka.group = lookupEnvDefault("KAFKA_GROUP", "storage")
	cfg.kafka.topics = strings.Split(lookupEnvDefault("KAFKA_TOPICS", "messages"), ",")

	return cfg, nil
}

func lookupEnvDefault(key string, def string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return def
	}
	return value
}
