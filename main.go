package main

import (
	"log"
	"SEHO/internal/config"
	"SEHO/internal/music"
	"SEHO/internal/redis"
    "SEHO/internal/logging"
)

func main() {
    //Setup logger and cleanup
    cleanup := logging.SetupLogger()
	defer cleanup()

	cfg := config.LoadConfig()
	rdb := redis.NewClient(&redis.Options{
        Addr: cfg:RedisAddress,
        )}

	music.StartMonitoring(cfg.MusicDirectory, rdb, 10*time.Second)
}

