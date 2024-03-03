package mongo

import (
	"context"
	"fmt"
	"newm/internal/app/db"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDb() (*mongo.Database, error) {
	v := viper.New()
	v.AddConfigPath("internal/app/db/mongo")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	db, err := initDb(db.Config{
		Host:     v.GetString("host"),
		Port:     v.GetString("port"),
		User:     v.GetString("user"),
		Password: v.GetString("password"),
		DbName:   v.GetString("dbName"),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDb(c db.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.DbName))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(c.DbName), nil
}
