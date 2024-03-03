package redis

import (
	"context"
	"fmt"
	"newm/internal/app/db"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func OpenDb() (*redis.Client, error) {
	viper.AddConfigPath("internal/app/db/mongo")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	db, err := initDb(db.Config{
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		User:     viper.GetString("user"),
		Password: viper.GetString("password"),
		DbName:   viper.GetString("dbName"),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDb(c db.Config) (*redis.Client, error) {
	dbName, err := strconv.Atoi(c.DbName)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		Password: c.Password,
		DB:       dbName,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
