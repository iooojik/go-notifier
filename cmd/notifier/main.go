package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go-notifier/internal"
	"log"
	"strings"
)

const (
	lastGoVersionId = "last_go_version"
)

func readCfg() {
	viper.SetConfigFile(".cfg")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Config file changed:", e.Name)
		})
		viper.WatchConfig()
	}()
}

func main() {
	readCfg()

	goVersion, err := internal.GetLatestGoVersion(viper.GetString("GO_VERSION_SOURCE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("REDIS_DSN"),
	})

	savedVersion, err := rdb.Get(context.Background(), lastGoVersionId).Result()
	if !errors.Is(err, redis.Nil) && err != nil {
		log.Fatalln(err)
	}

	fmt.Println("last go version", savedVersion, "got version:", goVersion)

	if savedVersion == goVersion {
		return
	}

	err = internal.SendMessage(goVersion, viper.GetString("GO_VERSION_SOURCE_URL"), viper.GetString("BOT_TOKEN"), strings.Split(viper.GetString("CHAT_IDS"), ","))
	if err != nil {
		log.Fatalln(err)
	}

	err = rdb.Set(context.Background(), lastGoVersionId, goVersion, 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("saved new go version", goVersion)
}
